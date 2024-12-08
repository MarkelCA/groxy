package main

import (
	"net"
)

func main() {
	src_addr := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	dst_addr := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8081,
	}

	// f := NewSecureForwarder(src_addr, dst_addr)
	f := NewForwarder(src_addr, dst_addr)
	f.Init()
}
