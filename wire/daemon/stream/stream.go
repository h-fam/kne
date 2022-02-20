package stream

import (
	"io"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"k8s.io/klog"
)

type Wire struct {
	node string
	intf string
}

type Stream struct {
	peer     Wire
	intf     string
	h        *pcap.Handle
	sendChan chan []byte

	mu     sync.Mutex
	closed bool
}

func New(intf string, w Wire) (*Stream, error) {
	h, err := pcap.OpenLive(intf, 10000, true, pcap.BlockForever)
	if err != nil {
		return nil, err
	}
	s := &Stream{
		h:    h,
		intf: intf,
		peer: w,
	}
	go s.startRecv()
	go s.startSend()
	return s, nil
}

func (s *Stream) Send(b []byte) error {
	if s.isClosed() {
		return io.EOF
	}
	s.sendChan <- b
}

func (s *Stream) startRecv() {
	for {
		b, ci, err := s.h.ReadPacketData()
		if err != nil {
			s.appendErr(err)
			return
		}
		p := gopacket.NewPacket(b, layers.LayerTypeEthernet, gopacket.Default)
		klog.V(100).Infof("%+v recieved %s: %s", s.wire, ci.Timestamp, p.String())
		if err := s.peer.Write(b); err != nil {
			s.appendErr(err)
			return
		}
	}
}

func (s *Stream) startSend() {
	for {
		wp := <-s.sendChan
		p := gopacket.NewPacket(wp.Bytes, layers.LayerTypeEthernet, gopacket.Default)
		klog.V(100).Infof("%+v writing %s: %s", s.wire, ci.Timestamp, p.String())
		if err := s.h.WritePacketData(p.Bytes); err != nil {
			s.appendErr(err)
			return
		}
	}
}

func (s *Stream) isClosed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.closed
}
