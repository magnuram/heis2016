//This is a test main
package main

import (
	. "./src/config"
	"./src/elev"
	//. "./src/queue"
	. "./src/driver"
	"log"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//backup_recovery()

	const elevDelay = 50 * time.Millisecond
	const openDoor = 1000 * time.Millisecond

	//_____________init hardware

	log.Println("Main: \t Start in main")
	buttonChannel := make(chan elev.ElevButton, 10)
	lightChannel := make(chan elev.ElevLight)
	//elevinfoChannel := make(chan driver.ElevInfo)
	motorChannel := make(chan int)
	floorChannel := make(chan int)
	defer close(buttonChannel)
	defer close(lightChannel)
	defer close(motorChannel)
	defer close(floorChannel)

	if err := elev.ElevInit(buttonChannel, lightChannel, motorChannel, floorChannel, elevDelay); err != nil {
		log.Println("ERROR -> Main: \t Hardware init failure")
		log.Fatal(err)
	} else {
		log.Println("Hardware init complete")
	}
	var floor = <-floorChannel
	//var floor  driver.ElevInfo
	//var elevator config.ElevInfo
	//elevator.Dir = elevator_type.Stop
	var light elev.ElevLight

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

	buttonLightOn := func(flr int, btn int) {
		light.Active = true
		light.Type = btn
		light.Floor = flr
		lightChannel <- light
	}

	buttonLightOff := func(flr int, btn int) {
		light.Active = false
		light.Type = btn
		light.Floor = flr
		lightChannel <- light
	}

	gotoFloor := func(flr int) { //Makes the elevator go to the floor

		if floor > flr {
			motorChannel <- DOWN
		} else if floor < flr {
			motorChannel <- UP
		}
		for floor != flr {
			floor = <-floorChannel
		}
		doorCheck()
	}


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


	//go buttonLightOn(floorChannel, buttonChannel)
	//ElevButton{Type: BUTTON_STOP}
	//driver.ElevLight{Type: INDICATOR_DOOR, Active: True}

	//if (motorChannel <- UP || motorChannel <- DOWN) {
	//log.Println("Que: ",que(button , floor , buttonChannelMatrix))
	for {
		//log.Println("Floorchannel: \n" ,floor) //0 -> 3
		//fmt.Printf("ButtonChannel: %v \n" ,<- buttonChannel) //{0 0}

		select {
		case btn := <-buttonChannel:
			switch btn.Type {
			//-----------------------------------------------External button
			case 0: //-------------------UP BUTTON
				switch btn.Floor {
				case 0: //1.etg
					buttonLightOn(0, BUTTON_CALL_UP)
					gotoFloor(0)
					buttonLightOff(0, BUTTON_CALL_UP)
				case 1: //2.etg
					//buttonLightOn(1, BUTTON_CALL_UP)
					gotoFloor(1)
					buttonLightOff(1, BUTTON_CALL_UP)
				case 2: //3.etg
					//buttonLightOn(2, BUTTON_CALL_UP)
					gotoFloor(2)
					buttonLightOff(2, BUTTON_CALL_UP)
				case 3: //4.etg
					//buttonLightOn(3, BUTTON_CALL_UP)
					gotoFloor(3)
					buttonLightOff(3, BUTTON_CALL_UP)
				}
			case 1: //----------------------DOWN Button
				switch btn.Floor {
				case 0: //1.etg
					//buttonLightOn(0, BUTTON_CALL_DOWN)
					gotoFloor(0)
					buttonLightOff(0, BUTTON_CALL_DOWN)
				case 1: //2.etg
					//buttonLightOn(1, BUTTON_CALL_DOWN)
					gotoFloor(1)
					buttonLightOff(1, BUTTON_CALL_DOWN)
				case 2: //3.etg
					//buttonLightOn(2, BUTTON_CALL_DOWN)
					gotoFloor(2)
					buttonLightOff(2, BUTTON_CALL_DOWN)
				case 3: //4.etg
					//buttonLightOn(3, BUTTON_CALL_DOWN)
					gotoFloor(3)
					buttonLightOff(3, BUTTON_CALL_DOWN)
				}
				//---------------------------------------------Local button
			case 2:
				switch btn.Floor {
				case 0: //1.etg
					//buttonLightOn(0, BUTTON_COMMAND)
					gotoFloor(0)
					buttonLightOff(0, BUTTON_COMMAND)
				case 1: //2.etg
					//buttonLightOn(1, BUTTON_COMMAND)
					gotoFloor(1)
					buttonLightOff(1, BUTTON_COMMAND)
				case 2: //3.etg
					//buttonLightOn(2, BUTTON_COMMAND)
					gotoFloor(2)
					buttonLightOff(2, BUTTON_COMMAND)
				case 3: //4.etg
					//buttonLightOn(3, BUTTON_COMMAND)
					gotoFloor(3)
					buttonLightOff(3, BUTTON_COMMAND)
				}
			default:
				log.Printf("Fail button")

			} //switch

		} //select

	} //for

} //main
