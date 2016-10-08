package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func run(tcpConn *net.TCPConn) {
	var serverLine string

	log.Printf("TCP connect: remote=%v local=%v", tcpConn.RemoteAddr(), tcpConn.LocalAddr())

	if buf, err := ioutil.ReadAll(tcpConn); err != nil {
		log.Fatalf("TCP read error: %v", err)
	} else {
		serverLine = string(buf)
	}

	fmt.Printf("client: %25v -> %-25v\n", tcpConn.LocalAddr(), tcpConn.RemoteAddr())
	fmt.Printf("server: %v\n", serverLine)

	if err := tcpConn.Close(); err != nil {
		log.Fatalf("TCP close error: %v", err)
	}
}

func main() {
	var options struct {
		host string
		port string
	}

	flag.StringVar(&options.port, "connect-port", "1337", "Listen port")
	options.host = flag.Arg(0)

	var addr = net.JoinHostPort(options.host, options.port)

	if tcpAddr, err := net.ResolveTCPAddr("tcp", addr); err != nil {
		log.Fatalf("net.ResolveTCPAddr %v: %v", addr, err)
	} else if tcpConn, err := net.DialTCP("tcp", nil, tcpAddr); err != nil {
		log.Fatalf("net.DialTCP %v: %v", tcpAddr, err)
	} else {
		log.Printf("run addr=%v: ...", tcpAddr)

		run(tcpConn)
	}
}
