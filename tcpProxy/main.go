package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	defer server.Close()

	remoteAddrs := []string{"tcp-echo.fly.dev:5001", "bad-addr.testing:9000"}

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		connectionManager(conn, remoteAddrs)
	}
}

func randomTarget(remoteAddrs []string) string {
	randNum := rand.Intn(len(remoteAddrs))
	fmt.Println(randNum)
	return remoteAddrs[randNum]

}

func connectionManager(conn net.Conn, remoteAddrs []string) {
	target, err := net.Dial("tcp", randomTarget(remoteAddrs))
	if err != nil {
		fmt.Println("bad connection, trying again")
		connectionManager(conn, remoteAddrs)
		return
	}
	go proxyConnection(conn, target)
}

func proxyConnection(conn net.Conn, targetConn backend) {
	defer conn.Close()

	go io.Copy(conn, targetConn)
	io.Copy(targetConn, conn)
}

func addTarget(conn net.TCPConn) {}

type backend net.Conn
