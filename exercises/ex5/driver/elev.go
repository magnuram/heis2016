package driver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "elev.h"
#include "elev.c"
*/

import (
	"C"
	"log"
)

const (
	N_FLOORS  = 4 //Number of floors, Hardware-depentent
	N_BUTTONS = 3 //Number of buttons/lamps on a per-floor basis
)

//
const (
	BUTTON_CALL_UP   = 0
	BUTTON_CALL_DOWN = 1
	BUTTON_COMMAND   = 2
)

//Motor commands
const (
	DIRN_DOWN = -1
	DIRN_STOP = 0
	DIRN_UP   = 1
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

var lampChannelMatrix = [N_FLOORS][N_BUTTONS]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var buttonChannelMatrix = [N_FLOORS][N_BUTTONS]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func ElevInit(buttonChannel chan<- ElevButton, lightChannel <-chan ElevLight) error {
	//init the hardware
	if err := IoInit(); err != nil {
		log.Println("in ElevInit():\t IoInit() ERROR")
		return err
	}

	//Set at floor button lamps off
	for Type := BUTTON_CALL_UP; Type <= BUTTON_COMMAND; Type++ {
		for floor := 0; floor < N_FLOORS; floor++ {
			IoClearBit(lampChannelMatrix[floor][Type])
		}
	}
	IoClearBit(LIGHT_DOOR_OPEN)
	IoClearBit(LIGHT_STOP)

}

func ElevSetMotorDirection() {

}

func ElevSetButtonLamp(floor int, button int, value bool) {

}

func ElevSetFloorIndicator() {

}

func ElevSetDoorOpenLamp() {

}

func ElevSetStopLamp(value bool) {
	if value {
		IoSetBit(LIGHT_STOP)
	} else {
		IoClearBit(LIGHT_STOP)
	}
}

func ElevGetButtonSignal() {

}

func ElevGetFloorSensorSignal() {

}

func ElevGetStopSignal() bool {
	return IoReadBit(STOP)
}

func ElevGetObstructuionSignal() bool {
	return IoReadBit(OBSTRUCTION)
}
