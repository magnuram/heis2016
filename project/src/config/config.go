package config

const NumbrOfFloors int = 4

//Motor commands
const UP = 1
const STOP = 0
const DOWN = -1

//Enumerator
const (
	ButtonCallUp = iota
	ButtonCallDown
	ButtonLocal
	SensorFloor
	IndicatorFloor
	ButtonStop   
	SensorObst   
	INDICATOR_DOOR
)

//Buttons 
const(
	UpButton = iota
	DownButton 
	LocalButton 
)

//Floors
const(
FirstFloor = iota
SecondFloor
ThirdFloor
FourthFloor
)