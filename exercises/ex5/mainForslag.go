package main

import (
	//"fmt"
	"log"
)
import "./driverForslag"

func main() {

	if err := driverForslag.ElevInit(); err != nil {
		log.Println("ERROR -> Main: \t Hardware init failure")
		log.Fatal(err)
	} else {
		log.Println("Hardware init complete")
	}
	driverForslag.ElevSetFloorIndicator(0)
	/*
		fmt.Println("Press STOP button to stop elevator and exit program.\n")

		driverForslag.ElevSetMotorDirection(driverForslag.UP)

		for {

			select {
			case floor := driverForslag.ElevGetFloorSensorSignal():
				switch floor {
				case 3:
					driverForslag.ElevSetFloorIndicator(3)
					driverForslag.ElevSetMotorDirection(driverForslag.DOWN)
				case 2:
					driverForslag.ElevSetFloorIndicator(2)
				case 1:
					driverForslag.ElevSetFloorIndicator(1)
				case 0:
					driverForslag.ElevSetFloorIndicator(0)
					driverForslag.ElevSetMotorDirection(driverForslag.UP)
				default:
					fmt.Println("Fail value", floor)
				}
			}

			if driverForslag.ElevGetStopSignal() {
				driverForslag.ElevSetMotorDirection(driverForslag.STOP)

			}

			/*
				//enkel test up og ned
				if driverForslag.ElevGetFloorSensorSignal() == driverForslag.N_FLOORS -1 {
					driverForslag.ElevSetMotorDirection(driverForslag.DOWN)
				} else if driverForslag.ElevGetFloorSensorSignal() == 0 {
					driverForslag.ElevSetMotorDirection(driverForslag.UP)
				}

				if driverForslag.ElevGetStopSignal() {
					driverForslag.ElevSetMotorDirection(driverForslag.STOP)

				}
	*/

}
