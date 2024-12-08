package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type Forwarder struct {
	src    net.TCPAddr
	dst    net.TCPAddr
	client net.Conn
	remote net.Conn
}

func NewForwarder(src_addr, dst_addr net.TCPAddr) Forwarder {
	log.Println(fmt.Sprintf("Forwarding from port %d to %s", src_addr.Port, dst_addr.String()))
	ln, err := net.Listen("tcp", src_addr.String())
	if err != nil {
		log.Fatal(err)
	}
	client, err := ln.Accept()

	if err != nil {
		log.Fatal("Error accepting connection")
	}

	remote, err := net.Dial("tcp", dst_addr.String())
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal("Error accepting connection")
	}
	return Forwarder{src_addr, dst_addr, client, remote}
}

func (f Forwarder) Init() {
	for {
		message, err := bufio.NewReader(f.client).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		f.remote.Write([]byte(newmessage + "\n"))
	}
}
