package udp

import (
	"net"
	"strings"
	"encoding/json"
	"fmt"

	)
const (
	PORT = ":20009"
	)


const (
	IP = 100
	ELEVATOR_PASSES_FLOOR = 101
	ELEVATOR_STOPS_AT_FLOOR = 102
	ELEVATOR_STARTS_RUNNING = 105
	ELEVATOR_ADDED_TO_NTWK = 103
	ELEVATOR_REMOVED_FROM_NTWK = 104
	ACK = 106

	MASTER = -1 // If MASTER is target, the message is meant for the current master
	)

type Message struct {
	Source int
	Id int
	Floor int
	Target int
	Checksum int
}

func (m *Message) Calc_checksum() int {
	return m.Source % 7 + m.Id % 7 + m.Floor % 7 + m.Target % 7 
}

func UDPSender(channel chan Message) {
	broadcastAddr := []string{"129.241.187.255", PORT}
	broadcastUDP, _ := net.ResolveUDPAddr("udp", strings.Join(broadcastAddr, ""))
	broadcastConn, _ := net.DialUDP("udp", nil, broadcastUDP)
	defer broadcastConn.Close()
	for {
		buf, err := json.Marshal(<- channel)
		if err == nil {
			broadcastConn.Write(buf)
		}
	}
}

func UDPListener(channel chan Message) {
	UDPReceiveAddr, err := net.ResolveUDPAddr("udp", PORT);
	if err != nil { fmt.Println(err) }

	UDPConn, err := net.ListenUDP("udp", UDPReceiveAddr);
	if err != nil { fmt.Println(err) }
	defer UDPConn.Close()

	buf := make([]byte, 2048)
	trimmed_buf := make([]byte, 1)
	var received_message Message

	for {
		n, _, _ := UDPConn.ReadFromUDP(buf)
		trimmed_buf = buf[:n]
		err := json.Unmarshal(trimmed_buf, &received_message)
		if err == nil {
			channel <- received_message
		}
	}
}