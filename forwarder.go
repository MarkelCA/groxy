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
	key    []byte
}

func xor(input []byte, key []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}
	return output
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

	key := []byte("password")

	return Forwarder{src_addr, dst_addr, client, remote, key}
}

func (f Forwarder) Init() {
	for {
		message, err := bufio.NewReader(f.client).ReadString('\n')
		message = string(xor([]byte(message), f.key))
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		f.remote.Write(xor([]byte(newmessage+"\n"), f.key))
	}
}
