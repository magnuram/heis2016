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
	IndicatorDoor
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

//lights ON or OFF
const(
	ON	= true
	OFF = false
)


// Colours for printing to console
const Col0 = "\x1b[30;1m" // Dark grey
const ColR = "\x1b[31;1m" // Red
const ColG = "\x1b[32;1m" // Green
const ColY = "\x1b[33;1m" // Yellow
const ColB = "\x1b[34;1m" // Blue
const ColM = "\x1b[35;1m" // Magenta
const ColC = "\x1b[36;1m" // Cyan
const ColW = "\x1b[37;1m" // White
const ColN = "\x1b[0m"    // Grey (neutral)