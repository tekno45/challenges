package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {

	remoteAddrs, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var conf Config

	err = json.Unmarshal(remoteAddrs, &conf)
	var wg sync.WaitGroup
	for _, v := range conf.Apps {
		for _, p := range v.Ports {
			port := strconv.Itoa(p)
			listenerAddr := fmt.Sprintf("0.0.0.0:%s", port)
			go ProxyListener(listenerAddr, v.Targets)
			wg.Add(1)
		}
	}
	wg.Wait()

	//[]string{"tcp-echo.fly.dev:5001", "bad-addr.testing:9000"}
}

type App struct {
	Name    string   `json:"Name"`
	Ports   []int    `json:"Ports"`
	Targets []string `json:"Targets"`
}

type Config struct {
	Apps []App `json:"Apps"`
}

func ProxyListener(listenerAddr string, remoteAddrs []string) {
	listener, err := net.Listen("tcp", listenerAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		connectionManager(conn, remoteAddrs)
	}
}

func randomTarget(remoteAddrs []string) string {
	randNum := rand.Intn(len(remoteAddrs))
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

type backend net.Conn
