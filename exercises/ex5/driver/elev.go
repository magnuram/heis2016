package driver

import (
	//"C"
	"log"
	"math"
	"time"
)

const (
	N_FLOORS  = 4 //Number of floors, Hardware-depentent
	N_BUTTONS = 3 //Number of buttons/lamps on a per-floor basis
)

//
const (
	BUTTON_CALL_UP = iota
	BUTTON_CALL_DOWN
	BUTTON_COMMAND
	SENSOR_FLOOR
	INDICATOR_FLOOR
	BUTTON_STOP
	SENSOR_OBST
	INDICATOR_DOOR
)

var ButtonType = []string{
	"BUTTON_CALL_UP",
	"BUTTON_CALL_DOWN",
	"BUTTON_COMMAND",
	"SENSOR_FLOOR",
	"INDICATOR_FLOOR",
	"BUTTON_STOP",
	"SENSOR_OBST",
	"INDICATOR_DOOR",
}

var MotorCmd = [3]string{
	"DOWN",
	"STOP",
	"UP",
}

const maxSpeed int = 14
const elevStopDelay = 50 * time.Millisecond

//Motor commands

const (
	UP   = 1
	STOP = 0
	DOWN = -1
)

type ElevButton struct {
	Type  int
	Floor int
}

type ElevLight struct {
	Type   int
	Floor  int
	Active bool
}

var lampChannelMatrix = [N_FLOORS][3]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var buttonChannelMatrix = [N_FLOORS][3]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func ElevInit(buttonchannel chan<- ElevButton, lightChannel <-chan ElevLight, motorChannel chan int, floorChannel chan<- int, elevDelay time.Duration) error {
	//init the hardware

	if err := IoInit(); err != nil {
		log.Println("ElevInit():\t IoInit() ERROR")
		return err
	}

	clearAlllights()

	go lightCheck(lightChannel)

	go elevSetMotorDirection(motorChannel)

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

func readInput(buttonchannel chan<- ElevButton, elevDelay time.Duration) {
	inputMatrix := [N_FLOORS][3]bool{}
	var stopbtn bool = false
	for {
		for Type := BUTTON_CALL_UP; Type <= BUTTON_COMMAND; Type++ {
			for Floor := 0; Floor < N_FLOORS; Floor++ {
				tempbtn := IoReadBit(buttonChannelMatrix[Floor][Type])
				if tempbtn { // button been pressed
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
				buttonchannel <- ElevButton{Type: BUTTON_STOP}
			}
		} else {
			stopbtn = false
		}
		time.Sleep(elevDelay)
	}

}
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

func lightCheck(lightChannel <-chan ElevLight) {
	var cmd ElevLight
	for {
		select {
		case cmd = <-lightChannel:
			switch cmd.Type {
			case BUTTON_STOP:
				if cmd.Active {
					IoSetBit(LIGHT_STOP)
				} else {
					IoClearBit(LIGHT_STOP)
				}
			case BUTTON_CALL_UP, BUTTON_CALL_DOWN, BUTTON_COMMAND:
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

func elevSetMotorDirection(motorChannel <-chan int) {
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

func setFloorIndicator(floor int) {
	if floor >= N_FLOORS {
		floor = N_FLOORS - 1
		log.Println("Elev: \t Tried to set the light indicator to the one over", N_FLOORS-1)
	} else if floor < 0 {
		floor = 0
		log.Println("Elev: \t Tried to set the light indicator to under 0")
	}
	if floor&0x02 > 0 { //floor&0x02 > 0 | bool((floor & 0x02) != 0)
		IoSetBit(LIGHT_FLOOR_IND1)
	} else {
		IoClearBit(LIGHT_FLOOR_IND1)
	}
	if floor&0x01 > 0 { // floor&0x01 > 0 | bool((floor & 0x01) != 0)
		IoSetBit(LIGHT_FLOOR_IND2)
	} else {
		IoClearBit(LIGHT_FLOOR_IND2)
	}
}

func getFloorSensorSignal() int {
	if IoReadBit(SENSOR_FLOOR1) {
		return 0
	} else if IoReadBit(SENSOR_FLOOR2) {
		return 1
	} else if IoReadBit(SENSOR_FLOOR3) {
		return 2
	} else if IoReadBit(SENSOR_FLOOR4) {
		return 3
	} else {
		return -1
	}
}

func ElevGetStopSignal() bool {
	return IoReadBit(STOP_BUTTON)
}

func ElevGetObstructuionSignal() bool { //not going to use
	return IoReadBit(OBSTRUCTION)
}

//---------------------------------SubFunctions-----------------------------------------//
func clearAlllights() {
	//Set at floor button lamps off
	for Type := BUTTON_CALL_UP; Type <= BUTTON_COMMAND; Type++ { //buttonCallUp = 1 Button_command = 3
		for Floor := 0; Floor < N_FLOORS; Floor++ {
			IoClearBit(lampChannelMatrix[Floor][Type])
		}
	}
	IoClearBit(LIGHT_DOOR_OPEN)

	IoClearBit(LIGHT_STOP)

}
