package network


import (
	"fmt"
	"net"
	"strconv"
)

var LocAddr *net.UDPAddr //Local address
var BroAddr *net.UDPAddr //Broadcast address