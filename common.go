package jproxy

import (
	"io"
	"log"
	"net"

	"golang.org/x/net/proxy"
)

func RunProxy(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Failed to create listener:", err)
	}

	log.Println("Jproxy listening on ", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	socks5Conn, err := proxy.SOCKS5("tcp", "", nil, proxy.Direct)
	if err != nil {
		log.Println("Failed to create proxy: ", err)
		return
	}

	target, err := socks5Conn.Dial("tcp", conn.RemoteAddr().String())
	if err != nil {
		log.Println("Failed to connect to target: ", err)
		return
	}
	defer target.Close()

	done := make(chan struct{})
	go func() {
		io.Copy(target, conn)
		close(done)
	}()

	go func() {
		io.Copy(conn, target)
		close(done)
	}()

	<-done
}
