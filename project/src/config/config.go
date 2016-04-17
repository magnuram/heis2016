package config



const N_FLOORS int = 4

//Motor commands
const UP = 1
const STOP = 0
const DOWN = -1

//Enumerator
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


