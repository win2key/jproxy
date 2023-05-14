package jproxy

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
)

const (
	socksVer5       = 5
	socksCmdConnect = 1
)

func RunProxy(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Failed to create listener:", err)
	}

	log.Println("Jproxy listening on", address)

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

	// Read the handshake message
	buf := make([]byte, 2)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		log.Println("failed to read handshake:", err)
		return
	}
	if buf[0] != socksVer5 {
		log.Println("invalid SOCKS version:", buf[0])
		return
	}

	// Read the authentication methods
	nmethods := int(buf[1])
	methods := make([]byte, nmethods)
	_, err = io.ReadFull(conn, methods)
	if err != nil {
		log.Println("failed to read methods:", err)
		return
	}

	// Send the handshake response
	_, err = conn.Write([]byte{socksVer5, 0})
	if err != nil {
		log.Println("failed to write handshake response:", err)
		return
	}

	// Read the request
	buf = make([]byte, 4)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		log.Println("failed to read request:", err)
		return
	}
	if buf[0] != socksVer5 {
		log.Println("invalid SOCKS version:", buf[0])
		return
	}
	if buf[1] != socksCmdConnect {
		log.Println("unsupported command:", buf[1])
		return
	}

	// Read the target address and port
	addr, err := readAddr(conn)
	if err != nil {
		log.Println("failed to read address:", err)
		return
	}
	target, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println("failed to connect to target:", err)
		return
	}
	defer target.Close()

	// Send the response
	_, err = conn.Write([]byte{socksVer5, 0, 0, 1, 0, 0, 0, 0, 0, 0})
	if err != nil {
		log.Println("failed to write response:", err)
		return
	}

	// Relay data
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, _ = io.Copy(target, conn)
	}()
	go func() {
		defer wg.Done()
		_, _ = io.Copy(conn, target)
	}()
	wg.Wait()
}

func readAddr(conn net.Conn) (string, error) {
	buf := make([]byte, 4)
	_, err := io.ReadFull(conn, buf)
	if err != nil {
		return "", err
	}
	ip := net.IP(buf)
	buf = make([]byte, 2)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return "", err
	}
	port := binary.BigEndian.Uint16(buf)
	return net.JoinHostPort(ip.String(), strconv.Itoa(int(port))), nil
}
