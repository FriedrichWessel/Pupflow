package main

import (
	"net"
)

func startNetworking(saddr string, objects <-chan []byte) {
	addr, e := net.ResolveUDPAddr("udp4", saddr)
	if e != nil {
		panic(e)
	}

	c, e := net.DialUDP("udp4", nil, addr)
	if e != nil {
		panic(e)
	}

	for object := range objects {
		c.Write(append(object, '\n'))
	}
}


