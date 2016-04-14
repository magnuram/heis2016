package config

//import "name"

const (
	N_FLOORS  int = 4 //Number of floors, Hardware-depentent
	N_BUTTONS     = 3 //Number of buttons/lamps on a per-floor basis
)

const (
	STOP_DIR = 0
	UP_DIR   = 1
	DOWN_DIR = -1
)

const (
	BUTTON_CALL_UP   = iota //0
	BUTTON_CALL_DOWN        // 1
	BUTTON_COMMAND          //2
	SENSOR_FLOOR            //3
	INDICATOR_FLOOR         //4
	BUTTON_STOP             //5
	SENSOR_OBST             //6
	INDICATOR_DOOR          //7
)

//STATES
const (
	INIT = iota
	IDLE
	MOVING
	DOOROPEN
	STOPS
)

//Motor Command
const (
	UP   = 1
	STOP = 0
	DOWN = -1
)

type ELEVINFO struct {
	State     int
	PrevFloor int
	Direction int
	ReqUp     [N_FLOORS]int
	ReqDown   [N_FLOORS]int
	ReqLocal  [N_FLOORS]int
}

var Elinf = ELEVINFO{}
