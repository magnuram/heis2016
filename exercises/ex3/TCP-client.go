
package main

import (
	
	"fmt"
	"time"
	"net"
	"bufio"
	"os"
	"runtime"
	//"strconv"

)

const (
	server = "129.241.187.23"
	port_fixed = "34933"
	port_0_delimited = "33546" // Alle meldinger ma avsluttes med \0
)

func les_og_skriv(reader *bufio.Reader) {
	melding := make([]byte, 1024)

	a, err := reader.Read(melding)

	if a != 1024 {
		fmt.Println("Meldingen er enten kortere eller lenger enn 1024 Byte")
	}
	if err != nil {
		//fmt.Println("Error ved lesing av melding")
		//fmt.Fprintln(os.Stderr, "read size: " + strconv.Itoa(a))
		//fmt.Fprintln(os.Stderr, "Fixed reader error: " + err.Error())
		return
	}
	fmt.Println(string(melding[:a]))

}


func connect_to_TCP(port string, callback func(*bufio.Reader)){
	conn, err := net.Dial("tcp", net.JoinHostPort(server, port_fixed))
	if err != nil {
        fmt.Fprintln(os.Stderr, "connection error on " + server + ":" + port)
        fmt.Fprintln(os.Stderr, "Connection join error: " + err.Error())
        return
	}
	defer conn.Close()

	prat_med_server(conn, callback)


}


func prat_med_server(conn net.Conn, callback func(*bufio.Reader)) {
	reader := bufio.NewReaderSize(conn, 1024)
	callback(reader)
	time.Sleep(1000*time.Millisecond)
	fmt.Println("Tekst som skal sendes: ")
	tekst, _ := reader.ReadString('\n')
	// Send to socket
	fmt.Println(conn, tekst + "\n")
}




func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("********TCP-kommunikasjon**********")

	connect_to_TCP(port_fixed, les_og_skriv)

	time.Sleep(1000*time.Millisecond)


}








	//_, err := socket.Read(buffer[:])

	// Connect to this socket
/*
	
	conn, _ := net.Dial("tcp", net.JoinHostPort(server, port_fixed))

	

	if err != nil {
		fmt.Println("Error connecting to TCP server")
	}
	
	addr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Println("Failed to resolve address for: " + port_fixed)
	}

	listener, err := net.ListenTCP("tcp", addr)

	time.Sleep(1000*time.Millisecond)

}
*/


	/*
	for {
		// Read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Tekst som skal sendes: ")
		tekst, _ := reader.ReadString('\n')
		// Send to socket
		fmt.Println(conn, tekst + "\n")
		// Listen for reply
		melding, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message from server: "+melding)
		}
		time.Sleep(1000*time.Millisecond)
}
*/
/*
	conn, err := net.Dial("tcp", net.JoinHostPort(server, port_fixed))
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	time.Sleep(1000*time.Millisecond)

}

*/

