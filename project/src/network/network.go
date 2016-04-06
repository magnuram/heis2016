package network


import (
	"fmt"
	"net"
	"strconv"
)

var LocAddr *net.UDPAddr //Local address
var BroAddr *net.UDPAddr //Broadcast address

type Udp_message struct {
	Raddr  string //if receiving raddr=senders address, if sending raddr should be set to "broadcast" or an ip:port
	Data   []byte //TODO: implement another encoding, strings are meh
	Length int    //length of received data, in #bytes // N/A for sending
}

func UdpInit(localListenPort, broadcastListenPort, msgSize int, sendChan, recieveChan chan UdpMessage)(err error){
	//generating broadcast address
	BroAddr, err = net.ResolveUDPAddr("udp4", "255.255.255.255:"+strconv.Itoa(broadcastListenPort))
	if err != nil{ return err }
}	