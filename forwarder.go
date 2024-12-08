package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type Forwarder struct {
	src_host net.IP
	src_port uint16
	dst_host net.IP
	dst_port uint16

	conn net.Conn
}

func NewForwarder(src_host net.IP, src_port uint16, dst_host net.IP, dst_port uint16) Forwarder {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", src_port))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Listening on port %d", src_port))
	conn, err := ln.Accept()

	if err != nil {
		log.Fatal("Error accepting connection")
	}

	return Forwarder{src_host, src_port, dst_host, dst_port, conn}
}

func (f Forwarder) Init() {
	for {
		message, err := bufio.NewReader(f.conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		f.conn.Write([]byte(newmessage + "\n"))
	}
}
