// Tror dette er eneste fullførte for ex3
//This is the UDPsender.go file for ex3

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

	//net.UDPConn("")
	//set up send-socket
	remote_addr, _ := net.ResolveUDPAddr("udp", "129.241.187.255:20018") //129.241.187.255 using 15 because of workspace
	socket_send, err := net.DialUDP("udp", nil, remote_addr)
	CheckError(err)

	//set up a listen-socket
	port, _ := net.ResolveUDPAddr("udp", ":20018")
	socket_listen, err := net.ListenUDP("udp", port)
	CheckError(err)

	//close sockets when done
	defer socket_listen.Close()
	defer socket_send.Close()

	for {
		//Send message
		msg := "hei server"
		socket_send.Write([]byte(msg))

		//Listen to message
		var buffer [1024]byte //64
		length, addr, err := socket_listen.ReadFromUDP(buffer[:])
		log.Println(length)
		log.Println(addr)
		log.Println(err)
		log.Println(string(buffer[:]), "\n")
		time.Sleep(1 * time.Second)

	}

}
