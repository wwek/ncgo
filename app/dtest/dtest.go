package dtest

import (
	"net"
)

type Target struct {
	Net   string
	Taddr *net.UDPAddr
}

func (t *Target) Tconn() (*net.UDPConn, error) {
	return net.DialUDP(t.Net, nil, t.Taddr)
}

func Duc(net string, addr *net.UDPAddr) (*net.UDPConn, error) {
	return net.DialUDP()
}
