//This is a test main
package main

import (
	//"fmt"
	"log"
	"runtime"
	"time"
	"./driver"
	//"./config"
	//"encoding/binary"
	//"net"
	//"os/exec"
	//"os"
	//"os/signal"
)



const (
 UP   = driver.UP
 STOP = driver.STOP
 DOWN = driver.DOWN
)


const (
	BUTTON_CALL_UP 		= driver.BUTTON_CALL_UP 	//0  
	BUTTON_CALL_DOWN 	= driver.BUTTON_CALL_DOWN	// 1
	BUTTON_COMMAND  	= driver.BUTTON_COMMAND		//2
	SENSOR_FLOOR     	= driver.SENSOR_FLOOR   	//3
	INDICATOR_FLOOR 	= driver.INDICATOR_FLOOR	//4
	BUTTON_STOP  		= driver.BUTTON_STOP		//5
	SENSOR_OBST  		= driver.SENSOR_OBST		//6
	INDICATOR_DOOR 		= driver.INDICATOR_DOOR		//7
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	


	//backup_recovery()



	const elevDelay = 50 * time.Millisecond
	const openDoor = 3000 * time.Millisecond

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
			motorChannel <- STOP 			// Stops any initial elevator movement
			light.Active = true 			// Makes "door open" light for all elevators active
			light.Type = INDICATOR_DOOR		// Describes 
			lightChannel <- light 			
			time.Sleep(openDoor)			
			light.Active = false			
			lightChannel <- light 			
		}

	buttonLightRed := func (flr int) {
		light.Active = true
		light.Type = BUTTON_CALL_DOWN
		light.Floor = flr
		lightChannel <- light
	}
	buttonLightBlank := func(flr int){
		light.Active = false
		light.Type = BUTTON_CALL_DOWN
		light.Floor = flr
		lightChannel <- light
	}
	
			//ElevButton{Type: BUTTON_STOP}
			//driver.ElevLight{Type: INDICATOR_DOOR, Active: True}

	//if (motorChannel <- UP || motorChannel <- DOWN) {

	for {
		
		//fmt.Printf("Floorchannel: %v \n" ,<-floorChannel) //0 -> 3
		//fmt.Printf("ButtonChannel: %v \n" ,<- buttonChannel) //{0 0}	

		select {
		case btn := <-buttonChannel:
			switch btn.Type{
			case 0,1:			//external
				log.Println("meaow")

			case 2:				//Localbutton
				if btn.Floor == 0{
					if floor > 0 {
					motorChannel <- DOWN
				} else if floor < 0 {
					motorChannel <- UP
				}
				for floor != 0 {floor =<- floorChannel}
				buttonLightBlank(0)	
				doorCheck()
					}else if btn.Floor == 1{
						if floor > 1 {
					motorChannel <- DOWN
					} else if floor < 1 {
					motorChannel <- UP
					}				
					for floor != 1{floor =<- floorChannel}
					buttonLightBlank(1)
					doorCheck()

						}else if btn.Floor == 2{
							if floor > 2 {
						motorChannel <- DOWN
						} else if floor < 2 {
						motorChannel <- UP
						}
						for floor != 2{floor =<- floorChannel}
						buttonLightBlank(2)
						doorCheck()

							}else if btn.Floor == 3{
								if floor > 3 {
								motorChannel <- DOWN
								} else if floor < 3 {
								motorChannel <- UP
								}
								for floor != 3{floor =<- floorChannel}
								buttonLightBlank(3)
								doorCheck()

							}
			
			default:
			log.Printf("Fail button")

			}//switch
	
		}//select
		
	}//for

}//main



