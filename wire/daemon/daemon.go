package main

import (
	"flag"
	"net"
	"os"

	"github.com/google/kne/wire/daemon/cni"
	"k8s.io/klog/v2"
)

var (
	addr = flag.String("port", ":51111", "Default listen address for service")
)

func main() {
	flag.Parse()
	klog.InitFlags(flag.CommandLine)
	s, err := cni.New(cni.Config{})
	if err != nil {
		klog.Fatalf("Failed to create stream server: %v", err)
	}
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		klog.Fatalf("Failed to listen on %s: %v", addr, err)
	}
	if err := s.Serve(lis); err != nil {
		klog.Errorf("Server exited with error: %v", err)
		os.Exit(1)
	}
}
