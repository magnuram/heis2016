package driver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "elev.h"
#include "elev.c"
*/

import "C"

const (
	N_FLOORS  = 4 //Number of floors, Hardware-depentent
	N_BUTTONS = 3 //Number of buttons/lamps on a per-floor basis
)

type Elev_button_type_t int

const (
	BUTTON_CALL_UP   Elev_button_type_t = 0
	BUTTON_CALL_DOWN Elev_button_type_t = 1
	BUTTON_COMMAND   Elev_button_type_t = 2
)

type Elev_motor_direction_t int

const (
	DIRN_DOWN Elev_motor_direction_t = -1
	DIRN_STOP Elev_motor_direction_t = 0
	DIRN_UP   Elev_motor_direction_t = 1
)

var lamp_channel_matrix = [N_FLOORS][N_BUTTONS]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var button_channel_matrix = [N_FLOORS][N_BUTTONS]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func ElevInit()int{
	//init the hardware
	if !=IoInit(){
		return nil
	}

	//Set at floor button lamps off
	for floor := 0; i < N_FLOORS; floor++ {
		IoClearBit(lamp_channel_matrix[floor])
	}

}

func ElevSetMotorDirection() {
	
}

func ElevSetButtonLamp() {
	
}

func ElevSetFloorIndicator() {
	
}

func ElevSetDoorOpenLamp() {
	
}

func ElevSetStopLamp() {
	
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