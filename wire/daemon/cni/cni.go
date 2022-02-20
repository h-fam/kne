package cni

import (
	"context"
	"net"
	"path/filepath"

	topologyclientv1 "github.com/google/kne/api/clientset/v1beta1"
	wpb "github.com/google/kne/proto/wire"
	"github.com/google/kne/wire/daemon/stream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"
)

type Config struct {
	GRPCOpts []grpc.ServerOption
}

func restConfig() (*rest.Config, error) {
	klog.Infof("Trying in-cluster configuration")
	rCfg, err := rest.InClusterConfig()
	if err != nil {
		kubecfg := filepath.Join(".kube", "config")
		if home := homedir.HomeDir(); home != "" {
			kubecfg = filepath.Join(home, kubecfg)
		}
		klog.Infof("Falling back to kubeconfig: %q", kubecfg)
		rCfg, err = clientcmd.BuildConfigFromFlags("", kubecfg)
		if err != nil {
			return nil, err
		}
	}
	return rCfg, nil
}

type Server struct {
	config  Config
	rCfg    *rest.Config
	s       *grpc.Server
	lis     net.Listener
	tClient topologyclientv1.Interface
	kClient kubernetes.Interface
	streams map[stream.Wire]stream.Stream
}

// New returns a new Server for managing the wirestream RPC service.
// The stream service enables the local CNI to create veth to grpc wires.
// The controller is also response to connecting to or receiving
// connections from remote controllers via the wiresteam protocol.
func New(cfg Config) (*Server, error) {
	rCfg, err := restConfig()
	if err != nil {
		return nil, err
	}
	kClient, err := kubernetes.NewForConfig(rCfg)
	if err != nil {
		return nil, err
	}
	tClient, err := topologyclientv1.NewForConfig(rCfg)
	if err != nil {
		return nil, err
	}

	s := &Server{
		config:  cfg,
		rCfg:    rCfg,
		kClient: kClient,
		tClient: tClient,
		s:       grpc.NewServer(cfg.GRPCOpts...),
		streams: map[stream.Wire]stream.Stream{},
	}
	wpb.RegisterLocalServer(s.s, s)
	wpb.RegisterRemoteServer(s.s, s)
	reflection.Register(s.s)
	return s, nil
}

func (s *Server) Get(ctx context.Context, pod *wpb.PodQuery) (*wpb.Pod, error) {
	klog.Infof("Retrieving %s's metadata from K8s...", pod.Name)

	result, err := s.getPod(ctx, pod.Name, pod.KubeNs)
	if err != nil {
		klog.Errorf("Failed to read pod %s from K8s", pod.Name)
		return nil, err
	}

	remoteLinks, found, err := unstructured.NestedSlice(result.Object, "spec", "links")
	if err != nil || !found || remoteLinks == nil {
		klog.Errorf("Could not find 'Link' array in pod's spec")
		return nil, err
	}

	links := make([]*wpb.Link, len(remoteLinks))
	for i := range links {
		remoteLink, ok := remoteLinks[i].(map[string]interface{})
		if !ok {
			klog.Errorf("Unrecognised 'Link' structure")
			return nil, err
		}
		newLink := &wpb.Link{}
		newLink.PeerPod, _, _ = unstructured.NestedString(remoteLink, "peer_pod")
		newLink.PeerIntf, _, _ = unstructured.NestedString(remoteLink, "peer_intf")
		newLink.LocalIntf, _, _ = unstructured.NestedString(remoteLink, "local_intf")
		newLink.LocalIp, _, _ = unstructured.NestedString(remoteLink, "local_ip")
		newLink.PeerIp, _, _ = unstructured.NestedString(remoteLink, "peer_ip")
		newLink.Uid, _, _ = unstructured.NestedInt64(remoteLink, "uid")
		links[i] = newLink
	}

	srcIP, _, _ := unstructured.NestedString(result.Object, "status", "src_ip")
	netNs, _, _ := unstructured.NestedString(result.Object, "status", "net_ns")

	return &wpb.Pod{
		Name:   pod.Name,
		SrcIp:  srcIP,
		NetNs:  netNs,
		KubeNs: pod.KubeNs,
		Links:  links,
	}, nil
}

func (s *Server) SetAlive(ctx context.Context, pod *wpb.Pod) (*wpb.BoolResponse, error) {
	klog.Infof("Setting %s's SrcIp=%s and NetNs=%s", pod.Name, pod.SrcIp, pod.NetNs)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, err := s.getPod(ctx, pod.Name, pod.KubeNs)
		if err != nil {
			klog.Errorf("Failed to read pod %s from K8s", pod.Name)
			return err
		}

		if err = unstructured.SetNestedField(result.Object, pod.SrcIp, "status", "src_ip"); err != nil {
			klog.Errorf("Failed to update pod's src_ip")
		}

		if err = unstructured.SetNestedField(result.Object, pod.NetNs, "status", "net_ns"); err != nil {
			klog.Errorf("Failed to update pod's net_ns")
		}

		return s.updateStatus(ctx, result, pod.KubeNs)
	})
	if retryErr != nil {
		klog.Errorf("Failed to update pod %s alive status: %v", pod.Name, retryErr)
		return &wpb.BoolResponse{Response: false}, retryErr
	}
	return &wpb.BoolResponse{Response: true}, nil
}

func (s *Server) Skip(ctx context.Context, skip *wpb.SkipQuery) (*wpb.BoolResponse, error) {
	klog.Infof("Skipping of pod %s by pod %s", skip.Peer, skip.Pod)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, err := s.getPod(ctx, skip.Pod, skip.KubeNs)
		if err != nil {
			klog.Errorf("Failed to read pod %s from K8s", skip.Pod)
			return err
		}

		skipped, _, _ := unstructured.NestedSlice(result.Object, "status", "skipped")

		newSkipped := append(skipped, skip.Peer)

		if err := unstructured.SetNestedField(result.Object, newSkipped, "status", "skipped"); err != nil {
			klog.Errorf("Failed to updated skipped list")
			return err
		}

		return s.updateStatus(ctx, result, skip.KubeNs)
	})
	if retryErr != nil {
		klog.Errorf("Failed to update skip pod %s status", skip.Pod, retryErr)
		return &wpb.BoolResponse{Response: false}, retryErr
	}
	return &wpb.BoolResponse{Response: true}, nil
}

func (s *Server) SkipReverse(ctx context.Context, skip *wpb.SkipQuery) (*wpb.BoolResponse, error) {
	klog.Infof("Reverse-skipping of pod %s by pod %s", skip.Peer, skip.Pod)
	var podName string
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// setting the value for peer pod
		peerPod, err := s.getPod(ctx, skip.Peer, skip.KubeNs)
		if err != nil {
			klog.Errorf("Failed to read pod %s from K8s", skip.Pod)
			return err
		}
		podName = peerPod.GetName()
		// extracting peer pod's skipped list and adding this pod's name to it
		peerSkipped, _, _ := unstructured.NestedSlice(peerPod.Object, "status", "skipped")
		newPeerSkipped := append(peerSkipped, skip.Pod)
		klog.Infof("Updating peer skipped list")
		// updating peer pod's skipped list locally
		if err := unstructured.SetNestedField(peerPod.Object, newPeerSkipped, "status", "skipped"); err != nil {
			klog.Errorf("Failed to updated reverse-skipped list for peer pod %s", peerPod.GetName())
			return err
		}
		// sending peer pod's updates to k8s
		return s.updateStatus(ctx, peerPod, skip.KubeNs)
	})
	if retryErr != nil {
		klog.Errorf("Failed to update peer pod %s skipreverse status: %v", podName, retryErr)
		return &wpb.BoolResponse{Response: false}, retryErr
	}

	retryErr = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// setting the value for this pod
		thisPod, err := s.getPod(ctx, skip.Pod, skip.KubeNs)
		if err != nil {
			klog.Errorf("Failed to read pod %s from K8s", skip.Pod)
			return err
		}
		// extracting this pod's skipped list and removing peer pod's name from it
		thisSkipped, _, _ := unstructured.NestedSlice(thisPod.Object, "status", "skipped")
		newThisSkipped := make([]interface{}, 0)
		klog.Infof("THIS %v SKIPPED", thisSkipped)
		for _, el := range thisSkipped {
			elString, ok := el.(string)
			if ok {
				if elString != skip.Peer {
					klog.Errorf("Appending new element %s", elString)
					newThisSkipped = append(newThisSkipped, elString)
				}
			}
		}
		klog.Infof("NEW THIS %v SKIPPED", newThisSkipped)
		// updating this pod's skipped list locally
		if len(newThisSkipped) != 0 {
			if err := unstructured.SetNestedField(thisPod.Object, newThisSkipped, "status", "skipped"); err != nil {
				klog.Errorf("Failed to cleanup skipped list for pod %s", thisPod.GetName())
				return err
			}
			// sending this pod's updates to k8s
			return s.updateStatus(ctx, thisPod, skip.KubeNs)
		}
		return nil
	})
	if retryErr != nil {
		klog.Errorf("Failed to update this pod skipreverse status: %v", retryErr)
		return &wpb.BoolResponse{Response: false}, retryErr
	}
	return &wpb.BoolResponse{Response: true}, nil
}

func (s *Server) IsSkipped(ctx context.Context, skip *wpb.SkipQuery) (*wpb.BoolResponse, error) {
	klog.Infof("Checking if %s is skipped by %s", skip.Peer, skip.Pod)
	result, err := s.getPod(ctx, skip.Peer, skip.KubeNs)
	if err != nil {
		klog.Errorf("Failed to read pod %s from K8s", skip.Pod)
		return nil, err
	}
	skipped, _, _ := unstructured.NestedSlice(result.Object, "status", "skipped")
	for _, peer := range skipped {
		if skip.Pod == peer.(string) {
			return &wpb.BoolResponse{Response: true}, nil
		}
	}
	return &wpb.BoolResponse{Response: false}, nil
}

func (s *Server) Update(ctx context.Context, pod *wpb.RemotePod) (*wpb.BoolResponse, error) {
	if err := s.createOrUpdate(pod); err != nil {
		klog.Errorf("Failed to Update Vxlan")
		return &wpb.BoolResponse{Response: false}, nil
	}
	return &wpb.BoolResponse{Response: true}, nil
}

func (s *Server) createOrUpdate(pod *wpb.RemotePod) error {
	return nil
}

func (s *Server) Serve(lis net.Listener) error {
	klog.Infof("GRPC server has started on %s", lis.Addr().String())
	s.lis = lis
	return s.s.Serve(s.lis)
}

func (s *Server) Stop() {
	s.s.Stop()
}

func (s *Server) getPod(ctx context.Context, name, ns string) (*unstructured.Unstructured, error) {
	klog.Infof("Reading pod %s from K8s", name)
	return s.tClient.Topology(ns).Unstructured(ctx, name, metav1.GetOptions{})
}

func (s *Server) updateStatus(ctx context.Context, obj *unstructured.Unstructured, ns string) error {
	klog.Infof("Update pod status %s from K8s", obj.GetName())
	_, err := s.tClient.Topology(ns).Update(ctx, obj, metav1.UpdateOptions{})
	return err
}
