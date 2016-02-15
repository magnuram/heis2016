package main

import (
	
	"fmt"
	"time"
	"net"
	"bufio"
	"strings"

)


func main() {

	fmt.Println("TCP-server startet...\n")

	// Listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// Accept connection on port
	conn, _ := ln.Accept()

	// Run loop forever or until ctrl-c
	for {
		// Listens for message to prosess, ends with newline \n
		melding, _ := bufio.NewReader(conn).ReadString('\n')

		// Output-melding mottatt
		fmt.Print("Melding mottatt: ", string(melding))

		//Sample process for string received
		nymelding := strings.ToUpper(melding)

		// Send new string back to client
		conn.Write([]byte(nymelding + "\n"))
	}


	time.Sleep(1000*time.Millisecond)

}





	/*
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		//handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	// ...
	*/
