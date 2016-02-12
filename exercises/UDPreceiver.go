//This is the UDPresceiver.go file for ex3

package main

import (
	"log"
	"net"
	"time"
)

func CheckError(err error) { // error function
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:20021") //30000
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0
	for {
		msg := "HEI på deg"
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		CheckError(err)
		//var buffer [64]byte
		//length, addr, err := socket.ReadFromUDP(buffer[:])
		//log.Println(length)
		//log.Println(addr)
		//log.Println(err)
		//log.Println(string(buffer[:]), "\n")
		time.Sleep(1 * time.Second)
	}

}
