package main

import (
	"fmt"
	"net"
	"time"
)

const (
	host    = "129.241.187.255"
	udpPort = "20018"
	port_t  = "30000"
)

func udpReceive(port string) {
	buff := make([]byte, 1024)
	addr, _ := net.ResolveUDPAddr("udp", ":"+port)
	sock, _ := net.ListenUDP("udp", addr)
	for {
		_, _, err := sock.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buff[:]))
	}
}

func udpSend() {
	raddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, udpPort))
	if err != nil {
		fmt.Println("Failed to resolve address for: " + udpPort)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Println("EREREREROROOROROR")
	}

	go udpReceive(udpPort)
	for {
		time.Sleep(1000 * time.Millisecond)
		conn.Write([]byte("bananapancakes"))
		fmt.Println("Msg sent")
	}

}

func main() {
	udpSend()

}
