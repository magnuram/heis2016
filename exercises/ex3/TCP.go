//computer ip: 129.241.187.148
//port:34933 for fixed-size messages
//port: 33546 for 0-terminated messages

package main

import (
	
	"fmt"
	"time"
	"net"
	"bufio"
	"os"
	//"runtime"


)

const (
	server = "129.241.187.23"
	port_fixed = "34933"
	port_0_delimited = "33546" 
)

func main() {

	conn, _ := net.Dial("tcp", net.JoinHostPort(server, port_fixed))

	
		// Read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Tekst som skal sendes: ")
		tekst, _ := reader.ReadString('\n')
		// Send to socket
		fmt.Println(conn, tekst + "\n")
		// Listen for reply
		melding, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message from server: "+melding)
		

	time.Sleep(1000*time.Millisecond)
}
