//This is a test main
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

	const elevDelay = 50 * time.Millisecond
	const openDoor = 1000 * time.Millisecond


//------------------------Init Hardware
	log.Println("Main: \t Start in main")
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

	var floor = <-floorChannel
	var light elev.ElevLight

	doorCheck := func() {
		time.Sleep(elevDelay)       //Times Elevator at right place
		motorChannel <- STOP        // Stops any initial elevator movement
		light.Active = true         // Makes "door open" light for all elevators active
		light.Type = INDICATOR_DOOR // Describes
		lightChannel <- light
		time.Sleep(openDoor)
		light.Active = false
		lightChannel <- light
	}

	gotoFloor := func(flr int) { //Makes the elevator go to the floor

		if floor > flr {
			motorChannel <- DOWN
		} else if floor < flr {
			motorChannel <- UP
		}
		for floor != flr {			//Elevator mooving
			floor = <-floorChannel	
		}
		 doorCheck()
	}

	for {

		select {
			case btn := <-buttonChannel:
			lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: true}
			switch btn.Type {
			//-----------------------------------------------External button
				case UpButton: //------------------------------------------------UP Button
				switch btn.Floor {

					case FirstFloor: //1.etg
						log.Println("First floor: UP button pressed")
					 	gotoFloor(btn.Floor)
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

					case SecondFloor: //2.etg
						log.Println("Second floor: UP button pressed")
						gotoFloor(btn.Floor)
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

					case ThirdFloor: //3.etg
						log.Println("Third floor: UP button pressed")
						gotoFloor(btn.Floor)
						lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
				}

				case DownButton: //----------------------------------------------DOWN Button
					switch btn.Floor {

						case SecondFloor: //2.etg
							log.Println("Second floor: DOWN button pressed")
							gotoFloor(btn.Floor)
							lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

						case ThirdFloor: //3.etg
							log.Println("Third floor: DOWN button pressed")
							gotoFloor(btn.Floor)					
							lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

						case FourthFloor: //4.etg
							log.Println("Fourth floor: DOWN button pressed")
							gotoFloor(btn.Floor)
							lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
					}
					
				case LocalButton: //---------------------------------------------LOCAL Button
					switch btn.Floor {

						case FirstFloor: //1.etg
							log.Println("First floor: LOCAL button pressed")
							gotoFloor(btn.Floor)
							lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

						case SecondFloor: //2.etg
							log.Println("Second floor: LOCAL button pressed")
							gotoFloor(btn.Floor)
							lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

						case ThirdFloor: //3.etg
							log.Println("Third floor: LOCAL button pressed")
							gotoFloor(btn.Floor)
							lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}

						case FourthFloor: //4.etg
							log.Println("Fourth floor: LOCAL button pressed")
							gotoFloor(btn.Floor)
							lightChannel <- elev.ElevLight{Type: btn.Type, Floor: btn.Floor, Active: false}
					}
				default:
				log.Printf("Fail button")

			} //switch

		} //select

	} //for

}//main	
