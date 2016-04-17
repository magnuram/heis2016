package main

import (
	. "./src/config"
	"./src/elev"
	"log"
	"runtime"
	"time"
	"os"
	"os/signal"
)



func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Time delays
	const elevDelay = 50 * time.Millisecond
	const openDoor = 1000 * time.Millisecond


//------------------------Making channels
	log.Println(ColW,"Main started",ColG)
	buttonChannel := make(chan elev.ElevButton, 100)
	lightChannel := make(chan elev.ElevLight)
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
		log.Println(ColM,"Hardware init complete",ColG)
	}

	var floor = <-floorChannel

//----------Init simple kill channel
	killChan := make(chan os.Signal)
	signal.Notify(killChan, os.Interrupt)
	go func() {
		<-killChan
		motorChannel <- STOP
		log.Println(ColR,"\n----------   Elevator is killed    -------------",ColG)
		time.Sleep(100 * time.Millisecond)
		os.Exit(1)
	}()


	//Set the Open door light on and off
	doorCheck := func() {
		time.Sleep(elevDelay)       
		motorChannel <- STOP        
		lightChannel <- elev.ElevLight{Type: IndicatorDoor,  Active: ON} //Turns ON door indicator 

		time.Sleep(openDoor)  

		lightChannel <- elev.ElevLight{Type: IndicatorDoor,  Active: OFF} //Turns OFF door indicator

	}

	//Set the motor direction 
	gotoFloor := func(typ int, flr int) { //Makes the elevator go to the floor

		if floor > flr {
			motorChannel <- DOWN
		} else if floor < flr {
			motorChannel <- UP
		}
		for floor != flr {			//Elevator mooving
			floor = <-floorChannel	
		}
		 doorCheck()
		 lightChannel <- elev.ElevLight{Type: typ, Floor: flr, Active: OFF} //Turns OFF button light
	}

	for {
		//--------Checks which button pressed 
		select {
			case btn := <-buttonChannel:
			lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: ON} //Turns ON button light
			switch btn.Type {
			//_________________________________________External buttons________________________
				case UpButton: //-----------------------------------------UP Button
				switch btn.Floor {

					case FirstFloor: 
						log.Println("First floor:",ColY,"UP ",ColG,"button pressed")
					 	gotoFloor(btn.Type,btn.Floor)

					case SecondFloor: 
						log.Println("Second floor:",ColY," UP ",ColG,"button pressed")
						gotoFloor(btn.Type,btn.Floor)

					case ThirdFloor: 
						log.Println("Third floor:",ColY," UP ",ColG,"button pressed")
						gotoFloor(btn.Type,btn.Floor)
				}

				case DownButton: //--------------------------------------DOWN Button
					switch btn.Floor {

						case SecondFloor: 
							log.Println("Second floor:",ColC," DOWN ",ColG,"button pressed")
							gotoFloor(btn.Type,btn.Floor)

						case ThirdFloor: 
							log.Println("Third floor:",ColC," DOWN",ColG ,"button pressed")
							gotoFloor(btn.Type,btn.Floor)					

						case FourthFloor: 
							log.Println("Fourth floor:",ColC," DOWN ",ColG,"button pressed",ColG)
							gotoFloor(btn.Type,btn.Floor)
					}
			//__________________________________________LOCAL Button__________________________
				case LocalButton: 
					switch btn.Floor {

						case FirstFloor: 
							log.Println("First floor:",ColB," LOCAL ",ColG,"button pressed")
							gotoFloor(btn.Type,btn.Floor)

						case SecondFloor: 
							log.Println("Second floor:",ColB," LOCAL ",ColG,"button pressed")
							gotoFloor(btn.Type,btn.Floor)

						case ThirdFloor: 
							log.Println("Third floor:",ColB," LOCAL ",ColG,"button pressed")
							gotoFloor(btn.Type,btn.Floor)

						case FourthFloor: 
							log.Println("Fourth floor:",ColB," LOCAL ",ColG,"button pressed")
							gotoFloor(btn.Type,btn.Floor)
					}
				default:
				log.Printf(ColR,"Fail button",ColG)

			} //switch

		} //select

	} //for

}//main	
