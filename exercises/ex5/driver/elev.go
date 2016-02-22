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

type Elev_motor_dir_t int
const (
	DIRN_DOWN Elev_motor_dir_t  = -1
	DIRN_STOP  Elev_motor_dir_t = 0
	DIRN_UP   Elev_motor_dir_t  = 1
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

func ElevInit() error {
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

	return nil
}

func ElevSetMotorDirection(dirn Elev_motor_dir_t) {
 if dirn == 0 {
 	IoWriteAnalog(MOTOR,0)
 }else if dirn > 0{
 	IoClearBit(MOTORDIR)
 	IoWriteAnalog(MOTOR, 2800)
 }else if dirn < 0{
 	IoSetBit(MOTORDIR)
 	IoWriteAnalog(MOTOR,2800)
 }

}

func ElevSetButtonLamp(value bool) {
	if value {
		IoSetBit(LIGHT_DOOR_OPEN)
	} else{
		IoClearBit(LIGHT_DOOR_OPEN)
	}
}

func ElevSetFloorIndicator(floor int) {
	if floor >= N_FLOORS{
		floor = N_FLOORS - 1
		log.Println("Elev: \t Tried to set the light indicator to the one over", N_FLOORS-1)
	}else if floor < 0 {
		floor = 0
		log.Println("Elev: \t Tried to set the light indicator to under 0")
	}
	if bool((floor & 0x02) != 0){
		IoSetBit(LIGHT_FLOOR_IND1)
	} else{
		IoClearBit(LIGHT_FLOOR_IND1)
	}
	if bool((floor & 0x01) != 0){
		IoSetBit(LIGHT_FLOOR_IND2)
	}else{
		IoClearBit(LIGHT_FLOOR_IND2)
	}
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

func ElevGetFloorSensorSignal() int {
	if IoReadBit(SENSOR_FLOOR1){
		return 0
	}else if IoReadBit(SENSOR_FLOOR2){
		return 1
	}else if IoReadBit(SENSOR_FLOOR3){
		return 2
	}else if IoReadBit(SENSOR_FLOOR4){
		return 3
	}else{return -1}
}

func ElevGetStopSignal() bool {
	return IoReadBit(STOP)
}

func ElevGetObstructuionSignal() bool {
	return IoReadBit(OBSTRUCTION)
}
