package network

import (
    . "./src/udp"
    "time"
    "net"
  
    "fmt"
    )

var elev_timers map[int]*time.Timer

func broadCastIP(ID int, UDPsend chan Message) {
    for {
        UDPsend <- Message{Source: ID, Id: IP}
        time.Sleep(100 * time.Millisecond)
    }
}

func Network_manager(from_main chan Message, to_main chan Message) {
    addr, _ := net.InterfaceAddrs()
    self_id := int(addr[1].String()[12] - '0') * 100 + int(addr[1].String()[13] - '0') * 10 + int(addr[1].String()[14] - '0')

    UDPsend := make(chan Message, 100)
    UDPreceive := make(chan Message, 100)

    
    go broadCastIP(self_id, UDPsend)
    go UDPListener(UDPreceive)
    go UDPSender(UDPsend)

    elev_timers = make(map[int]*time.Timer)
    
    for {
        select {
            case message := <- UDPreceive:
                _, present := elev_timers[message.Source]

                if message.Id == IP {
                    if message.Source != self_id {
                        if present {
                            elev_timers[message.Source].Reset(time.Second)
                        } else {
                            elev_timers[message.Source] = time.AfterFunc(time.Second, func() { remove_elev(message.Source, from_main) } )
                            to_main <- Message{Source: message.Source, Id: ELEVATOR_ADDED_TO_NTWK}
                        }
                    }
                    break
                }

                if message.Checksum != message.Calc_checksum() {
                    fmt.Println("CHECKSUM ERROR", message.Checksum, " vs ", message.Calc_checksum()) 
                    break
                }

                to_main <- message

            case message := <- from_main:
                message.Checksum = message.Calc_checksum()
                UDPsend <- message
        }

    }
}

func remove_elev(id int, to_main chan Message) {
    delete(elev_timers, id)
    to_main <- Message{Source: id, Id: ELEVATOR_REMOVED_FROM_NTWK}
}