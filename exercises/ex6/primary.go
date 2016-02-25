package main

import (
	"log"
	"net"
	"time"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func primary(start int) {
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:20063") //port 15 e føkked => 63
	errorHandler(err)

	udpBroadcast, err := net.DialUDP("udp", nil, udpAddr)
	errorHandler(err)

	msg := make([]byte, 1)

	for i := start; ; i++ { //Telleren
		log.Println(i)
		msg[0] = byte(i)
		udpBroadcast.Write(msg) //sender verdien via udp
		time.Sleep(200 * time.Millisecond)
	}

	udpBroadcast.Close()
}

func main() {
	primary(1) //start verdi 1
}
