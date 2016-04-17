//This is a test main
package main

import (
	//"builtin"
	. "./src/config"
	"./src/elev"
	//"./src/driver"
	// "./src/network"
	// "./src/cost"
	"log"
	"runtime"
	"time"
	"os"
	"os/signal"
	//"strconv"
	//"fmt"
	//"math/rand"
)

//const debug = false

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//new
//	const connectionAttempsLimit = 10
    //r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const elevDelay = 50 * time.Millisecond
	const openDoor = 1000 * time.Millisecond
	//const ackTimeout = 500 * time.Millisecond
/*
	//new
	var knownElevators = make(map[string]*Elevator) //key = IPadr
	var activeElevators = make(map[string]bool)     //key = IPadr
	const iAmAliveTickTime = 100 * time.Millisecond
	const iAmAliveLimit = 3*iAmAliveTickTime + 10*time.Millisecond
	var externalOrderMatrix [N_FLOORS][2]ElevOrder

	//var orderTimeout = 5*time.Second + time.Duration(r.Intn(2000))*time.Millisecond
	const doorWaitTime = 3000 * time.Millisecond
*/


//------------------------Init Hardware
	log.Println("Main: \t Start in main")
	buttonChannel := make(chan elev.ElevButton, 100)
	lightChannel := make(chan elev.ElevLight)
	//elevinfoChannel := make(chan driver.ElevInfo)
	motorChannel := make(chan int)
	floorChannel := make(chan int)
	defer close(buttonChannel)
	defer close(lightChannel)
	defer close(motorChannel)
	defer close(floorChannel)

//------------------------Init Hardware
	if err := elev.ElevInit(buttonChannel, lightChannel, motorChannel, floorChannel, elevDelay); err != nil {
		log.Println("ERROR -> Main: \t Hardware init failure")
		log.Fatal(err)
	} else {
		log.Println("Hardware init complete")
	}


//----------Init simple kill channel
	killChan := make(chan os.Signal)
	signal.Notify(killChan, os.Interrupt)
	go func() {
		<-killChan
		motorChannel <- STOP
		log.Println("\n----------   Elevator is killed    -------------")
		time.Sleep(100 * time.Millisecond)
		os.Exit(1)
	}()

/*

//--------Init network
	receiveOrderChannel := make(chan ElevOrderMessage, 5)
	

	sendOrderChannel := make(chan ElevOrderMessage)
	
	
	receiveRestoreChannel := make(chan ElevRestoreMessage, 5)
	
	//sendRestoreChannel := make(chan ElevRestoreMessage)
	
	localIP, err := initNetwork(connectionAttempsLimit, receiveOrderChannel, sendOrderChannel, receiveRestoreChannel, sendRestoreChannel)
	if err != nil {
		log.Println("MAIN:\t Network init failed")
		log.Fatal(err)
	} else if debug {
		log.Println("MAIN:\t Network init successful")
	}
*/

	//var button = <-buttonChannel
	var floor = <-floorChannel
	var light elev.ElevLight
/*
//-----Initialise state------
	log.Println("MAIN:    Sending out a request after my previous state")
	sendRestoreChannel <- ElevRestoreMessage{
		AskerIP: localIP,
		State:   ElevState{},
		Event:   EvRequestingState,
	}
	//knownElevators[localIP] = ResolveElevator(ElevState{LocalIP: localIP, LastFloor: <-floorChannel})
//	updateActiveElevators(knownElevators, activeElevators, localIP, iAmAliveLimit)
	log.Println("MAIN:    State init finished. Starting from floor:", floor)//knownElevators[localIP].State.LastFloor)

/*
//------INIT TImer
	//timeoutChannel := make(chan ExtendedElevOrder)
	doorTimer := time.NewTimer(time.Second)
	doorTimer.Stop()
*/
	
	//var q_type [] int
	//var q_floor [] int
	

	


	doorCheck := func() {
		time.Sleep(elevDelay)       //Sets it at right place
		motorChannel <- STOP        // Stops any initial elevator movement
		light.Active = true         // Makes "door open" light for all elevators active
		light.Type = INDICATOR_DOOR // Describes
		lightChannel <- light
		time.Sleep(openDoor)
		light.Active = false
		lightChannel <- light
	}

	gotoFloor := func(btn int, flr int) { //Makes the elevator go to the floor

		go func() {
			if flr != -1 {
				lightChannel <- elev.ElevLight{Type: btn, Floor: flr, Active: true}
			}else {
				lightChannel <- elev.ElevLight{Type: btn, Floor: flr, Active: false}
			}
		}()

		if floor > flr {
			motorChannel <- DOWN
		} else if floor < flr {
			motorChannel <- UP
		}
		for floor != flr {			//Elevator mooving
			floor = <-floorChannel
			
		}

		 doorCheck()
		//q_type = q_type[:len(q_type)-1]
		//q_floor = q_floor[:len(q_floor)-1]
	}
/*
	go func() {
		//for driver.IoReadBit(buttonChannel) {
			lightChannel <- elev.ElevLight{Type: button.Type, Floor: button.Floor, Active: true}
		//}
	}()
*/	
/*

		go func () {

				var btn =<-buttonChannel

				if len(q_type)> 1{
				
					gotoFloor(q_floor[1])
				if (q_type[1] == btn.Type && q_floor[1] == btn.Floor) {
					for i:=0; i<len(q_type); i++ {

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: true}
					break
				}
				} else {
					//	lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
					
				
				time.Sleep(elevDelay)
			}

			}
			
			
		}()
		*/


	//go buttonLightOn(floorChannel, buttonChannel)
	//ElevButton{Type: BUTTON_STOP}
	//driver.ElevLight{Type: INDICATOR_DOOR, Active: True}

	//if (motorChannel <- UP || motorChannel <- DOWN) {
	//log.Println("Que: ",que(button , floor , buttonChannelMatrix))
	for {

		select {
			case btn := <-buttonChannel:
			lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: true}
			switch btn.Type {
			//-----------------------------------------------External button
				case UpButton: //-------------------UP BUTTON
				switch btn.Floor {
					case FirstFloor: //1.etg
					//q_type = append(q_type,btn.Type)//,[btn.Floor]int)
					//q_floor = append(q_floor,btn.Floor)
					//log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
						log.Println("1.etg UP")
					 gotoFloor(btn.Type , btn.Floor)

				
					
					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

					case SecondFloor: //2.etg
					//q_type = append(q_type,btn.Type)//,[btn.Floor]int)
					//q_floor = append(q_floor,btn.Floor)
					//log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
					log.Println("2.etg UP")
						gotoFloor(btn.Type , btn.Floor)

					
				
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

					case ThirdFloor: //3.etg
					//q_type = append(q_type,btn.Type)//,[btn.Floor]int)
					//q_floor = append(q_floor,btn.Floor)
					//log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
					log.Println("3.etg UP")
						gotoFloor(btn.Type , btn.Floor)

					
					
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

					case FourthFloor: //4.etg ---------DOEST EXIST
					//q_type = append(q_type,btn.Type)//,[btn.Floor]int)
					//q_floor = append(q_floor,btn.Floor)
					//log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
					log.Println("4.etg UP")
						gotoFloor(btn.Type , btn.Floor)

					
					
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
			
				}
				case DownButton: //----------------------DOWN Button
				switch btn.Floor {
					case FirstFloor: //1.etg  ---------- DOESNT EXIST
					//q_type = append(q_type,btn.Type)//,[btn.Floor]int)
					//q_floor = append(q_floor,btn.Floor)
					//log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
					//log.Println("Quetype: ",q_type)
				//log.Println("Quefloor: ",q_floor)
					log.Println("1.etg DOWN")
						gotoFloor(btn.Type , btn.Floor)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

					case SecondFloor: //2.etg
					//q_type = append(q_type,btn.Type)//,[btn.Floor]int)
					//q_floor = append(q_floor,btn.Floor)

					//log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
				//	log.Println("Quetype: ",q_type)
				//	log.Println("Quefloor: ",q_floor)
					log.Println("2.etg DOWN")
						gotoFloor(btn.Type , btn.Floor)
					
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

					case ThirdFloor: //3.etg
					//q_type = append(q_type,btn.Type)//,[btn.Floor]int)
					//q_floor = append(q_floor,btn.Floor)
					//log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
					//log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
					log.Println("3.etg DOWN")
						gotoFloor(btn.Type , btn.Floor)
					
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

					case FourthFloor: //4.etg
					//q_type = append(q_type,btn.Type)//,[btn.Floor]int)
					//q_floor = append(q_floor,btn.Floor)
					//log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
				//	log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
					log.Println("4.etg DOWN")
						gotoFloor(btn.Type , btn.Floor)

						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
					}
					//---------------------------------------------Local button
				case LocalButton:
				switch btn.Floor {
					case FirstFloor: //1.etg
				//	q_type = append(q_type,btn.Type)//,[btn.Floor]int)
				//	q_floor = append(q_floor,btn.Floor)
						log.Println("1.etg local")
				//		log.Println("Quetype: ",q_type)
				//	log.Println("Quefloor: ",q_floor)
						 gotoFloor(btn.Type , btn.Floor)

						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
					case SecondFloor: //2.etg
				//	q_type = append(q_type,btn.Type)//,[btn.Floor]int)
				//	q_floor = append(q_floor,btn.Floor)
				//	log.Println("Quetype: ",q_type)
				//	log.Println("Quefloor: ",q_floor)
					log.Println("2.etg local")
						gotoFloor(btn.Type , btn.Floor)
						
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
					case ThirdFloor: //3.etg
				//	q_type = append(q_type,btn.Type)//,[btn.Floor]int)
				//	q_floor = append(q_floor,btn.Floor)
				//	log.Println("Quetype: ",q_type)
				//	log.Println("Quefloor: ",q_floor)
					log.Println("3.etg local")
						gotoFloor(btn.Type , btn.Floor)
						
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
					case FourthFloor: //4.etg
						log.Println("4.etg local")
					//	q_type = append(q_type,btn.Type)//,[btn.Floor]int)
					//q_floor = append(q_floor,btn.Floor)
					//	log.Println("Quetype: ",q_type)
					//log.Println("Quefloor: ",q_floor)
						gotoFloor(btn.Type , btn.Floor)
						
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
					}
				default:
				log.Printf("Fail button")

			} //switch

		} //select

	}
}	


 //main
/*
		if len(q_floor) > 0{
		flro := q_floor[0]

		if floor > flro {
			motorChannel <- DOWN
		} else if floor < flro {
			motorChannel <- UP
		}
		for floor != flro {
			floor = <-floorChannel
		}
	}
	*/
/*
select{
	//-------HARDWARE-------
		case button := <-buttonChannel:
			log.Println("MAIN:\t Received a", ButtonType[button.Type], "from floor", button.Floor, ".Number of activeElevators", len(activeElevators))
			switch button.Type {
			case BUTTON_CALL_UP, BUTTON_CALL_DOWN:
				if _, ok := activeElevators[localIP]; !ok {
					log.Println("MAIN:\t Can not accept new external order while offline!")
				} else {
					if assignedIP, err := cost.AssignNewOrder(knownElevators, activeElevators, externalOrderMatrix, button.Floor, button.Type); err != nil {
						log.Fatal(err)
					} else {
						sendOrderChannel <- ElevOrderMessage{
							Floor:      button.Floor,
							ButtonType: button.Type,
							AssignedTo: assignedIP,
							OriginIP:   localIP,
							SenderIP:   localIP,
							Event:      EvNewOrder,
						}
					}
				}
			case BUTTON_COMMAND:
				if !knownElevators[localIP].State.IsMoving && knownElevators[localIP].State.LastFloor == button.Floor {
					lightChannel <- elev.ElevLight{Type: INDICATOR_DOOR, Active: true}
					log.Println("MAIN:\t Opening doors")
					doorTimer.Reset(doorWaitTime)
					knownElevators[localIP].State.DoorIsOpen = true
					sendRestoreChannel <- ResolveBackupState(knownElevators[localIP], externalOrderMatrix)
				} else {
					printDebug("Added internal order to queue")
					knownElevators[localIP].SetInternalOrder(button.Floor)
					sendRestoreChannel <- ResolveBackupState(knownElevators[localIP], externalOrderMatrix)
					lightChannel <- elev.ElevLight{Type: button.Type, Floor: button.Floor, Active: true}
					if knownElevators[localIP].IsIdle() && !knownElevators[localIP].State.DoorIsOpen {
						doorTimer.Reset(0 * time.Millisecond)
					}
				}

			case BUTTON_STOP:
				motorChannel <- STOP
				lightChannel <- elev.ElevLight{Type: BUTTON_STOP, Active: true}
				fmt.Println("\n---------------------         SOMEBODY KILLED THIS ELEVATOR!     ---------------------")
				time.Sleep(200 * time.Millisecond)
				os.Exit(1)
			default:
				printDebug("Recived an ButtonType from the elev driver")
			}

		case floor := <-floorChannel:
			log.Println("MAIN:\t evFloorReached: ", floor)
			knownElevators[localIP].SetLastFloor(floor)
			if knownElevators[localIP].ResolveExtendedElevState(externalOrderMatrix).ShouldStop() {
				motorChannel <- STOP
				knownElevators[localIP].SetMoving(false)
				log.Println("MAIN:\t Opening doors")
				doorTimer.Reset(doorWaitTime)
				lightChannel <- elev.ElevLight{Type: INDICATOR_DOOR, Active: true}
				knownElevators[localIP].ClearInternalOrderAtCurrentFloor()
				lightChannel <- elev.ElevLight{Floor: floor, Type: BUTTON_COMMAND, Active: false}
				orders := knownElevators[localIP].ResolveExtendedElevState(externalOrderMatrix).FindExternalOrdersAtCurrentFloor()
				for _, o := range orders {
					externalOrderMatrix[o.Floor][o.Type].Status = NotActive
					externalOrderMatrix[o.Floor][o.Type].AssignedTo = ""
					externalOrderMatrix[o.Floor][o.Type].DeleteConfirmedBy()
					printDebug("Stoping timeoutTimer [Execution timeout] on order " + ButtonType[o.Type] + " on floor " + strconv.Itoa(o.Floor))
					externalOrderMatrix[o.Floor][o.Type].StopTimer()
					lightChannel <- elev.ElevLight{Floor: o.Floor, Type: o.Type, Active: false}
					externalOrderMatrix[o.Floor][o.Type].Timer = time.AfterFunc(ackTimeout, func() {
						log.Println("TIMEOUT:\t An orderDone was not ack´d by all activeElevators. Resending...")
						sendOrderChannel <- ElevOrderMessage{
							Floor:      o.Floor,
							ButtonType: o.Type,
							AssignedTo: o.Order.AssignedTo,
							OriginIP:   o.OriginIP,
							SenderIP:   localIP,
							Event:      EvOrderDone,
						}
					})
					printDebug("Sending orderDoneMessage on " + ButtonType[o.Type] + " on floor " + strconv.Itoa(o.Floor))
					sendOrderChannel <- ElevOrderMessage{
						Floor:      o.Floor,
						ButtonType: o.Type,
						AssignedTo: o.Order.AssignedTo,
						OriginIP:   o.OriginIP,
						SenderIP:   localIP,
						Event:      EvOrderDone,
					}
				}
			}
			sendRestoreChannel <- ResolveBackupState(knownElevators[localIP], externalOrderMatrix)
		
	}
*/
		//log.Println("Floorchannel: \n" ,floor) //0 -> 3
		//fmt.Printf("ButtonChannel: %v \n" ,<- buttonChannel) //{0 0}
	/*	
		select {
		case btn := <-buttonChannel:
		//	lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: true}
			switch btn.Type {
			//-----------------------------------------------External button
			case 0: //-------------------UP BUTTON
				switch btn.Floor {
				case 0: //1.etg
					gotoFloor(0)

				lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				case 1: //2.etg
					gotoFloor(1)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				case 2: //3.etg

					gotoFloor(2)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				case 3: //4.etg

					gotoFloor(3)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				}
			case 1: //----------------------DOWN Button
				switch btn.Floor {
				case 0: //1.etg

					gotoFloor(0)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				case 1: //2.etg

					gotoFloor(1)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				case 2: //3.etg

					gotoFloor(2)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				case 3: //4.etg

					gotoFloor(3)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				}
				//---------------------------------------------Local button
			case 2:
				switch btn.Floor {
				case 0: //1.etg

					gotoFloor(0)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				case 1: //2.etg

					gotoFloor(1)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				case 2: //3.etg

					gotoFloor(2)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				case 3: //4.etg

					gotoFloor(3)

					lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				}
			default:
				log.Printf("Fail button")

			} //switch

		} //select
		*/


/*
//new functions
func initNetwork(connectionAttempsLimit int, receiveOrderChannel, sendOrderChannel chan ElevOrderMessage, receiveRestoreChannel, sendRestoreChannel chan ElevRestoreMessage) (localIP string, err error) {
	for i := 0; i <= connectionAttempsLimit; i++ {
		localIP, err := network.Init(receiveOrderChannel, sendOrderChannel, receiveRestoreChannel, sendRestoreChannel)
		if err != nil {
			if i == 0 {
				log.Println("MAIN:\t Network init was not successful. Trying some more times")
			} else if i == connectionAttempsLimit {
				return "", err
			}
			time.Sleep(3 * time.Second)
		} else {
			return localIP, nil
		}
	}
	return "", nil
}


func updateActiveElevators(knownElevators map[string]*Elevator, activeElevators map[string]bool, localIP string, iAmAliveLimit time.Duration) {
	for key := range knownElevators {
		if time.Since(knownElevators[key].Time) > iAmAliveLimit {
			if activeElevators[key] == true {
				log.Printf("MAIN:\t Removed elevator %s in activeElevators\n", knownElevators[key].State.LocalIP)
				delete(activeElevators, key)
			}
		} else {
			if activeElevators[key] != true {
				activeElevators[key] = true
				log.Printf("MAIN:\t Added elevator %s in activeElevators\n", knownElevators[key].State.LocalIP)
			}
		}
	}
}


func printDebug(s string) {
	if debug {
		log.Println("MAIN:\t", s)
	}
}

/*
	elevGetButtonSignal := func(buttonChannel int, floorChannel int, buttonChannelMatrix[][] int) bool {
		if IoReadBit(buttonChannelMatrix[floor][button]) == true {
			return true
		} else {
			return false
		}
	}

	que := func(button int , floor int, buttonChannelMatrix [][] int) {

		var BtnChnl_UpList []int
		var BtnChnl_DownList []int
		var q []int
		//var buttonbool bool

		for elevGetButtonSignal(buttonChannel int, floorChannel int , buttonChannelMatrix int) == true {
			//var q [][][] int = IoReadBit(buttonChannelMatrix[floor][button])
			//currentFloor = ElevGetFloorSensorSignal()
			BtnChnl_UpList = buttonChannelMatrix[0] + buttonChannelMatrix[0] + buttonChannelMatrix[0] + buttonChannelMatrix[0]
			BtnChnl_DownList = buttonChannelMatrix[1] + buttonChannelMatrix[1] + buttonChannelMatrix[1] + buttonChannelMatrix[1]
			//BtnChnlMtrx_Command[] = buttonChannelMatrix{2,2,2,2}

			if IoReadBit(MOTORDIR) == true {
				q = BtnChnl_UpList + BtnChnl_DownList
			}
			if IoReadBit(MOTORDIR) == false {
				q = BtnChnl_DownList + BtnChnl_UpList
			}

		}

		return q
	}	
*/
	//btn := <-buttonChannel
			