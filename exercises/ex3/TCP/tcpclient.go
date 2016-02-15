package main

import (
	
	"fmt"
	"time"
	"net"
	"bufio"
	"os"

)



func main() {

	// Connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Tekst som skal sendes: ")
		tekst, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, tekst + "\n")
		melding, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+melding)
		}
	

	time.Sleep(1000*time.Millisecond)

}



