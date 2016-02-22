//This is a test main
package main

import (
	"fmt"
	"log"
	//"runtime"
)
import "./driver"

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//
	//buttonChannel := make(chan driver.ElevButton, 10)
	//lightChannel := make(chan driver.ElevLight)

	log.Println("Main: \t Start in main")

	driver.ElevInit()

	fmt.Println("Press STOP button to stop the elevator and exit the program\n")

	driver.ElevSetMotorDirection(driver.DIRN_UP)

	for {

		if driver.ElevGetFloorSensorSignal() == driver.N_FLOORS-1 {
			driver.ElevSetMotorDirection(driver.DIRN_DOWN)
		} else if driver.ElevGetFloorSensorSignal() == 0 {
			driver.ElevSetMotorDirection(driver.DIRN_UP)
		}

		if driver.ElevGetStopSignal() {
			driver.ElevSetMotorDirection(driver.DIRN_STOP)

		}
	}
}
