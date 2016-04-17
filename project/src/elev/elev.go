package elev

import (
	. "../channels"
	. "../config"
	. "../driver"
	"log"
	"math"
	"time"
)
//struct for elevator button: Type, Floor
type ElevButton struct {
	Type  int
	Floor int
}
//struct for elevator light: Type,Floor, Active
type ElevLight struct {
	Type   int
	Floor  int
	Active bool
}

const maxSpeed int = 14
const elevStopDelay = 50 * time.Millisecond

var lampChannelMatrix = [NumbrOfFloors][3]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var buttonChannelMatrix = [NumbrOfFloors][3]int{
	{BUTTON_UP1, BUTTON_DOWN1, ButtonLocal1},
	{BUTTON_UP2, BUTTON_DOWN2, ButtonLocal2},
	{BUTTON_UP3, BUTTON_DOWN3, ButtonLocal3},
	{BUTTON_UP4, BUTTON_DOWN4, ButtonLocal4},
}

func ElevInit(buttonchannel chan<- ElevButton, lightChannel <-chan ElevLight, motorChannel chan int, floorChannel chan<- int, elevDelay time.Duration) error {

	//--------------------------------------init the hardware
	if err := IoInit(); err != nil {
		log.Println("ElevInit():\t IoInit() ERROR")
		return err
	}

	clearAlllights()

	go lightController(lightChannel)

	go motorController(motorChannel)

	//Set elevator to a floor
	if getFloorSensorSignal() == -1 {
		motorChannel <- DOWN
		for {
			if getFloorSensorSignal() != -1 {
				motorChannel <- STOP
				break
			} else {
				time.Sleep(elevDelay)
			}
		}
	}

	go readInput(buttonchannel, elevDelay)

	go readFloorSensor(floorChannel, elevDelay)
	return nil
}

//Reads button pressed and sets it into channel
func readInput(buttonchannel chan<- ElevButton, elevDelay time.Duration) { 
	inputMatrix := [NumbrOfFloors][3]bool{}
	var stopbtn bool = false
	for {
		for Type := ButtonCallUp; Type <= ButtonLocal; Type++ {
			for Floor := 0; Floor < NumbrOfFloors; Floor++ {
				tempbtn := IoReadBit(buttonChannelMatrix[Floor][Type]) 
				if tempbtn {                                           
					if !inputMatrix[Floor][Type] {
						inputMatrix[Floor][Type] = true
						buttonchannel <- ElevButton{Type, Floor}
					}
				} else {
					inputMatrix[Floor][Type] = false
				}
			}
		}
		if IoReadBit(STOP_BUTTON) {
			if !stopbtn {
				stopbtn = true
				buttonchannel <- ElevButton{Type: ButtonStop}
			}
		} else {
			stopbtn = false
		}
		time.Sleep(elevDelay)
	}

}

//Drives the motor accordingly to what's in motorchannel
func motorController(motorChannel <-chan int) {
	for {
		select {
		case cmd := <-motorChannel:
			switch cmd {
			case STOP: 
				time.Sleep(elevStopDelay)
				IoWriteAnalog(MOTOR, 0)
			case UP: 
				IoClearBit(MOTORDIR)
				IoWriteAnalog(MOTOR, 200*int(math.Abs(float64(maxSpeed))))
			case DOWN: 
				IoSetBit(MOTORDIR)
				IoWriteAnalog(MOTOR, 200*int(math.Abs(float64(maxSpeed))))
			default:
				log.Println("Elev: \t invalid motor command: ", cmd)
			}
		}
	}
}

//Sets lights accordingly to what's in lightchannel
func lightController(lightChannel <-chan ElevLight) {
	var cmd ElevLight
	for {
		select {
		case cmd = <-lightChannel:
			switch cmd.Type {
			case ButtonStop: //-------------Doesn't use
				if cmd.Active {
					IoSetBit(LIGHT_STOP)
				} else {
					IoClearBit(LIGHT_STOP)
				}
			case ButtonCallUp, ButtonCallDown, ButtonLocal:
				if cmd.Active {
					IoSetBit(lampChannelMatrix[cmd.Floor][cmd.Type])
				} else {
					IoClearBit(lampChannelMatrix[cmd.Floor][cmd.Type])
				}
			case INDICATOR_DOOR:
				if cmd.Active {
					IoSetBit(LIGHT_DOOR_OPEN)
				} else {
					IoClearBit(LIGHT_DOOR_OPEN)
				}
			default:
				log.Println("Elev: \t You tried to light a non light item")
			}
		}

	}
}
//Get floor sensor signal and sends it to floorchannel
func readFloorSensor(floorChannel chan<- int, elevDelay time.Duration) { 
	var lastFloor int = -1
	for {
		tempFloor := getFloorSensorSignal()
		if (tempFloor != -1) && (tempFloor != lastFloor) {
			lastFloor = tempFloor
			setFloorIndicator(tempFloor)
			floorChannel <- tempFloor
		}
		time.Sleep(elevDelay)
	}
}

//Sets the floor indicator accordingly to what is in 
func setFloorIndicator(floor int) { 
	if floor >= NumbrOfFloors {
		floor = NumbrOfFloors - 1
		log.Println("Elev: \t Tried to set the light indicator to the one over", NumbrOfFloors-1)
	} else if floor < 0 {
		floor = 0
		log.Println("Elev: \t Tried to set the light indicator to under 0")
	}
	if floor&0x02 > 0 { 
		IoSetBit(LIGHT_FLOOR_IND1)
	} else {
		IoClearBit(LIGHT_FLOOR_IND1)
	}
	if floor&0x01 > 0 { 
		IoSetBit(LIGHT_FLOOR_IND2)
	} else {
		IoClearBit(LIGHT_FLOOR_IND2)
	}
}

//---------------------------------SubFunctions-----------------------------------------//
//Set at floor button lamps off
func clearAlllights() {
	for Type := ButtonCallUp; Type <= ButtonLocal; Type++ { 
		for Floor := 0; Floor < NumbrOfFloors; Floor++ {
			IoClearBit(lampChannelMatrix[Floor][Type])
		}
	}
	IoClearBit(LIGHT_DOOR_OPEN)

	IoClearBit(LIGHT_STOP)

}

//Read Floor sensors
func getFloorSensorSignal() int {
	if IoReadBit(SensorFloor1) {
		return 0 //0
	} else if IoReadBit(SensorFloor2) {
		return 1 //1
	} else if IoReadBit(SensorFloor3) {
		return 2 //2
	} else if IoReadBit(SensorFloor4) {
		return 3 //3
	} else {
		return -1
	}
}
