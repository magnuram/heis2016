package fsm

import (
	. "../config"
	. "../queue"
	"fmt"
)

var currentState int = Elinf.State
var direction int = 0
var currentFloor int = -1

func FsmInit(floorchannel <-chan int, motorChannel chan int) {
	var floor = <-floorChannel
	for {
		if floor != -1 {
			break
		}
		motorChannel <- DOWN
		if floor != 1 {
			currentFloor = floor
			motorChannel <- STOP
		}
	}
	currentState = IDLE
	fmt.Println("State:", currentState)
}

func FsmOrderExist() {
	var light elev.ElevLight
	switch currentState {
	case INIT:
		direction = QChooseDir(currentFloor, direction)
		motorChannel <- direction
		currentState = MOVING
		break
	case IDLE:
		direction = QChooseDir(currentFloor, direction)
		motorChannel <- direction
		currentState = MOVING
		break
	case MOVING:
		break
	case DOOROPEN:
		light.Active = false        // Makes "door open" light for all elevators active
		light.Type = INDICATOR_DOOR // Describes
		lightChannel <- light
		motorChannel <- direction

		currentState = MOVING
		break

	case STOP:
		break
	}
}
