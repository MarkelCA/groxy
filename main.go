package main

import (
	"net"
)

func main() {
	src_ip := net.ParseIP("127.0.0.1")
	dst_ip := net.ParseIP("127.0.0.1")

	f := NewForwarder(src_ip, 8080, dst_ip, 8081)
	f.Init()
}
