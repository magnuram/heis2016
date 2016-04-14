package driver

//import "name"

const (
	N_FLOORS  = 4 //Number of floors, Hardware-depentent
	N_BUTTONS = 3 //Number of buttons/lamps on a per-floor basis
)

const (
	BUTTON_CALL_UP = iota //0
	BUTTON_CALL_DOWN	// 1
	BUTTON_COMMAND  	//2
	SENSOR_FLOOR        //3
	INDICATOR_FLOOR 	//4
	BUTTON_STOP  		//5
	SENSOR_OBST  		//6
	INDICATOR_DOOR 		//7
)

//type  Dir int

const (
	 UP   	= 1
	 STOP  	= 0
	 DOWN  	= -1
)
/*
type ElevInfo struct {
	CurrentFloor 	int
	lol int
	Direction 		Dir
	//Requests 		[N_FLOORS][N_BUTTONS] int
	//State			ElevatorState 

}
*/
/*
type ButtonType int

const(
	ButtonLocal 		ButtonType = 0
	ButtonExternalUp 	ButtonType = 1
	ButtonExternalDown 	ButtonType =2
}
*/