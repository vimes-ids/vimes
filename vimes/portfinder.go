package vimes

import (
	"net"

	"github.com/cakturk/go-netstat/netstat"
)

type ListenPort struct {
	IP       net.IP
	Port     uint16
	Protocol string
}

type PortFinder struct {
	ports []ListenPort
}

func NewPortFinder() *PortFinder {
	return &PortFinder{}
}

func (list *PortFinder) GetGreeting() string {
	return "Hello world"
}

func (list *PortFinder) FindPorts() ([]ListenPort, error) {
	var filter = func(s *netstat.SockTabEntry) bool {
		return s.State == netstat.Listen
	}

	ports := []ListenPort{}

	callbacks := map[string]func(netstat.AcceptFn) ([]netstat.SockTabEntry, error){
		"tcp":  netstat.TCPSocks,
		"tcp6": netstat.TCP6Socks,
		"udp":  netstat.UDPSocks,
		"udp6": netstat.UDP6Socks,
	}

	for protocol, callback := range callbacks {
		sockets, _ := callback(filter)

		for _, socket := range sockets {
			ports = append(ports, ListenPort{
				IP:       socket.LocalAddr.IP,
				Port:     socket.LocalAddr.Port,
				Protocol: protocol,
			})
		}
	}

	return ports, nil
}
