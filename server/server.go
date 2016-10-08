package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func runAccept(tcpConn *net.TCPConn) {
	log.Printf("TCP accept: remote=%v local=%v", tcpConn.RemoteAddr(), tcpConn.LocalAddr())

	if _, err := fmt.Fprintf(tcpConn, "%25v -> %-25v", tcpConn.RemoteAddr(), tcpConn.LocalAddr()); err != nil {
		log.Printf("TCP write error: %v", err)
	}

	if err := tcpConn.Close(); err != nil {
		log.Printf("TCP close error: %v", err)
	}
}

func run(tcpListener *net.TCPListener) {
	log.Printf("Accepting TCP connections on %v...", tcpListener.Addr())

	for {
		if tcpConn, err := tcpListener.AcceptTCP(); err != nil {
			log.Fatalf("TCPListener.AcceptTCP: %v", err)
		} else {
			go runAccept(tcpConn)
		}
	}
}

func main() {
	var options struct {
		host string
		port string
	}

	flag.StringVar(&options.host, "listen-host", "0.0.0.0", "Listen host")
	flag.StringVar(&options.port, "listen-port", "1337", "Listen port")
	flag.Parse()

	var addr = net.JoinHostPort(options.host, options.port)

	if tcpAddr, err := net.ResolveTCPAddr("tcp", addr); err != nil {
		log.Fatalf("net.ResolveTCPAddr %v: %v", addr, err)
	} else if tcpListener, err := net.ListenTCP("tcp", tcpAddr); err != nil {
		log.Fatalf("net.ListenTCP %v: %v", tcpAddr, err)
	} else {
		log.Printf("run addr=%v: ...", tcpAddr)

		run(tcpListener)
	}
}
