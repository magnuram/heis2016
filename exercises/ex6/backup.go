package main

import (
	"encoding/binary"
	"log"
	"net"
	"os/exec"
	"time"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func primary(start int, udpBroadcast *net.UDPConn) {

	startNew()

	msg := make([]byte, 1)

	for i := start; ; i++ {
		log.Println(i)
		msg[0] = byte(i)
		udpBroadcast.Write(msg)
		time.Sleep(200 * time.Millisecond)
	}

	udpBroadcast.Close()
}

func startNew() {
	newBackup := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run backup.go") //setter opp en nytt vindu
	err := newBackup.Run()
	errorHandler(err)
}

func backup(udpListen *net.UDPConn) int {
	listenChn := make(chan int, 1)
	backupvalue := 0
	go listenFromUdp(listenChn, udpListen) //Kjører en goroutine som vil lese det som sendes fra primary
	for {
		select {
		case backupvalue = <-listenChn: // setter verdien i channel til backupverdien
			time.Sleep(100 * time.Millisecond)
			break
		case <-time.After(1 * time.Second): // Om tiden går lengre enn 1 sek vil den opprette en ny
			log.Println("The primary died, initiating backup")
			return backupvalue
		}
	}

}

func listenFromUdp(listenChn chan int, udpListen *net.UDPConn) {

	buff := make([]byte, 8)

	for {
		udpListen.ReadFromUDP(buff[:]) //Leser av buffer sendt på udp

		listenChn <- int(binary.LittleEndian.Uint64(buff)) //Konverterer bytearray til int og plasserer inn i channel
		time.Sleep(100 * time.Millisecond)
	}

}

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp", ":20063")
	errorHandler(err)

	udpListen, err := net.ListenUDP("udp", udpAddr)
	errorHandler(err)

	backupvalue := backup(udpListen) // Får ut backup verdi om d går lengre enn 1 sek

	udpListen.Close()

	udpAddr, err = net.ResolveUDPAddr("udp", "129.241.187.255:20063") // Oppretter ny oppkobling
	errorHandler(err)

	udpBroadcast, err := net.DialUDP("udp", nil, udpAddr)
	errorHandler(err)

	primary(backupvalue, udpBroadcast) //starter opp en ny "primary"

	//udpBroadcast.Close()

}
