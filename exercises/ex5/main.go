//This is a test main
package main



import (
	"fmt"
	"log"
	"runtime"
	"time"
)
import "./driver"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//

	const elevDelay = 50* time.Millisecond

	//_____________init hardware


	log.Println("Main: \t Start in main")
	buttonChannel := make(chan driver.ElevButton, 10)
	lightChannel := make(chan driver.ElevLight)
	motorChannel := make(chan int)
	floorChannel := make(chan int)


	if err := driver.ElevInit(buttonChannel, lightChannel, motorChannel, floorChannel, elevDelay); err != nil{
		log.Println("ERROR -> Main: \t Hardware init failure")
		log.Fatal(err)
	}else{
		log.Println("Hardware init complete")
	}

	fmt.Println("Press STOP button to stop the elevator and exit the program\n")

	// driver.ElevSetMotorDirection(driver.DIRN_UP)

	//for {

		
/*			Simple up and down test
		if driver.ElevGetFloorSensorSignal() == driver.N_FLOORS-1 {
			driver.ElevSetMotorDirection(driver.DIRN_DOWN)
		} else if driver.ElevGetFloorSensorSignal() == 0 {
			driver.ElevSetMotorDirection(driver.DIRN_UP)
		}
*/
		//if driver.ElevGetStopSignal() {
			//driver.ElevSetMotorDirection(driver.DIRN_STOP)

		//}
	//}
}