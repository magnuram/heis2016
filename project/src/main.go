//This is a test main
package main

import (
	//"fmt"
	"log"
	"runtime"
	"time"
	//"encoding/binary"
	//"net"
	//"os/exec"
	//"os"
	//"os/signal"
)
import  "./driver"


const (
	UP   = 1
	STOP = 0
	DOWN = -1
)


const (
	BUTTON_CALL_UP = iota //0  
	BUTTON_CALL_DOWN	// 1
	BUTTON_COMMAND  	//2
	SENSOR_FLOOR        //3
	INDICATOR_FLOOR 	//4
	BUTTON_STOP  		//5
	SENSOR_OBST  		//6
	INDICATOR_DOOR 		//7
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	


	//backup_recovery()



	const elevDelay = 50 * time.Millisecond
	const openDoor = 2000 * time.Millisecond

	//_____________init hardware

	log.Println("Main: \t Start in main")
	buttonChannel := make(chan driver.ElevButton, 10)
	lightChannel := make(chan driver.ElevLight)
	motorChannel := make(chan int)
	floorChannel := make(chan int)
	defer close(buttonChannel)
	defer close(lightChannel)
	defer close(motorChannel)
	defer close(floorChannel)

	if err := driver.ElevInit(buttonChannel, lightChannel, motorChannel, floorChannel, elevDelay); err != nil {
		log.Println("ERROR -> Main: \t Hardware init failure")
		log.Fatal(err)
	} else {
		log.Println("Hardware init complete")
	}
	var floor = <-floorChannel
	var light driver.ElevLight 

	doorCheck := func (){

			//time.Sleep(elevDelay)

			motorChannel <- STOP 			// Stops any initial elevator movement
			light.Active = true 			// Makes "door open" light for all elevators active
			light.Type = INDICATOR_DOOR		// Describes 
			lightChannel <- light 			
			
			time.Sleep(openDoor)			
			
			light.Active = false			
			lightChannel <- light 			
		}

/*
	stopCheck := func () {
				fmt.Printf("Stopper HEIS!!")
				motorChannel <- STOP
				light.Type =BUTTON_STOP
				light.Active = true
				lightChannel <- light
				time.Sleep(openDoor)
				light.Active = false
				lightChannel <- light
				motorChannel <- DOWN
				if floor != -1{
				motorChannel <-STOP
				}
				
	}
*/
	
			//ElevButton{Type: BUTTON_STOP}
			//driver.ElevLight{Type: INDICATOR_DOOR, Active: True}

	//if (motorChannel <- UP || motorChannel <- DOWN) {

	for {
		
		//fmt.Printf("Floorchannel: %v \n" ,<-floorChannel) //0 -> 3
		//fmt.Printf("ButtonChannel: %v \n" ,<- buttonChannel) //{0 0}	
		


		select {
		case btn := <-buttonChannel:
			switch btn.Floor{
			case 0:						//1.etg
				if driver.IoReadBit(driver.STOP_BUTTON) {
					motorChannel <- STOP  
					log.Println("Stoppknapp trykket")
					break
				}else if floor > 0 {
					motorChannel <- DOWN
					log.Println("Heis gar ned fra etasje ", floor, " til  0")
				} else if floor < 0 {
					motorChannel <- UP
					log.Println("Heis gar opp fra etasje ", floor, " til  0")
				}
				
				for floor != 0 {
			//		btn =<-buttonChannel}
					floor =<- floorChannel}
			//	if btn.Type == 5{stopCheck();break} }
				doorCheck()
			
			case 1: 					//2.etg
				if driver.IoReadBit(driver.STOP_BUTTON) {
					motorChannel <- STOP
					log.Println("Stoppknapp trykket")
					break
				}else if floor > 1 {
					motorChannel <- DOWN
					log.Println("Heis gar ned fra etasje ", floor, " til  1")
				} else if floor < 1 {
					motorChannel <- UP
					log.Println("Heis gar opp fra etasje ", floor, " til  1")
				}
				
			for floor != 1{ 
			//	btn =<-buttonChannel}
				floor =<- floorChannel}
			//	if btn.Type == 5 {stopCheck();break}}
			doorCheck()

			case 2:						 //3.etg
				if driver.IoReadBit(driver.STOP_BUTTON) {
					motorChannel <- STOP
					log.Println("Stoppknapp trykket")
					break
				}else if floor > 2 {
					motorChannel <- DOWN
					log.Println("Heis gar ned fra etasje ", floor, " til  2")
			} else if floor < 2 {
					motorChannel <- UP
					log.Println("Heis gar opp fra etasje ", floor, " til  2")
				}
			for floor != 2{ 
				floor =<- floorChannel}
			//	btn =<-buttonChannel
			//	if btn.Type == 5{stopCheck();break}}
			doorCheck()
			//motorChannel <- STOP
			
			//time.Sleep(openDoor) 

			case 3: 					//4.etg
				if driver.IoReadBit(driver.STOP_BUTTON) {
					motorChannel <- STOP
					log.Println("Stoppknapp trykket")
					break
				}else if floor > 3 {
					motorChannel <- DOWN
					log.Println("Heis gar ned fra etasje ", floor, " til  3")
			} else if floor < 3 {
					motorChannel <- UP
					log.Println("Heis gar opp fra etasje ", floor, " til  3")
				}
			for floor != 3{
				floor =<- floorChannel}
			//	btn =<-buttonChannel
			//	if btn.Type == 5{stopCheck();break}}
			doorCheck()
			//default:
			//fmt.Printf("Fail button")

			}//switch
	
		}//select

	}//for

}//main

	//} //







/*
   	//----------init monkey handle
   	killChannel := make(chan os.Signal)
   	signal.Notify(killChannel, os.Interrupt)
   	go func() {
   		<- killChannel
   		motorChannel <- 0
   		fmt.Println("KILLED ELEVATOR")
   		time.Sleep(100 * time.Millisecond)
   		os.Exit(1)
   	}()
   /*
   				Simple up and down test
   		if driver.ElevGetFloorSensorSignal() == driver.N_FLOORS-1 {
   			driver.ElevSetMotorDirection(driver.DIRN_DOWN)
   		} else if driver.ElevGetFloorSensorSignal() == 0 {
   			driver.ElevSetMotorDirection(driver.DIRN_UP)
   		}
*/
//if driver.ElevGetStopSignal() {
//			driver.ElevSetMotorDirection(driver.DIRN_STOP)

//		}
//}
