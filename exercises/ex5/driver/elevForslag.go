package driver

const (
	N_FLOORS  = 4 //Number of floors, Hardware-depentent
	N_BUTTONS = 3 //Number of buttons/lamps on a per-floor basis
)

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

const (
	MOTOR_SPEED = 2800
)

func Elev_init() {
	if err := IoInit(); err != nil {
		log.Println("ElevInit():\t IoInit() ERROR")
		return err
	}
}
func Elev_set_motor_direction(direction int) {
	if direction == 0 {
		IoWriteAnalog(MOTOR, 0)
	} else if direction > 0 {
		IoClearBit(MOTORDIR)
		IoWriteAnalog(MOTOR, MOTOR_SPEED)
	} else if direction < 0 {
		IoSetBit(MOTORDIR)
		IoWriteAnalog(MOTOR, MOTOR_SPEED)
	}
}

func Elev_set_button_lamp(button int, floor int, value bool) {
	if value {
		IoSetBit(lampChannelMatrix[floor][button])
	} else {
		IoClearBit(lampChannelMatrix[floor][button])
	}
}

func Elev_set_floor_indicator(floor int) {
	if floor&0x02 > 0 {
		IoSetBit(LIGHT_FLOOR_IND1)
	} else {
		IoClearBit(LIGHT_FLOOR_IND1)
	}

	if floor&0x01 > 0 {
		IoSetBit(LIGHT_FLOOR_IND2)
	} else {
		IoClearBit(LIGHT_FLOOR_IND2)
	}
}

func Elev_set_door_open_lamp(value bool) {
	if value {
		IoSetBit(LIGHT_DOOR_OPEN)
	} else {
		IoClearBit(LIGHT_DOOR_OPEN)
	}
}

func Elev_set_stop_lamp(value bool) {
	if value {
		IoSetBit(LIGHT_STOP)
	} else {
		IoClearBit(LIGHT_STOP)
	}
}

func Elev_get_button_signal(button int, floor int) bool {
	if IoReadBit(buttonChannelMatrix[floor][button]) == 1 {
		return true
	} else {
		return false
	}
}

func Elev_get_floor_sensor_signal() int {
	if IoReadBit(SENSOR_FLOOR1) == 1 {
		return 0
	} else if IoReadBit(SENSOR_FLOOR2) == 1 {
		return 1
	} else if IoReadBit(SENSOR_FLOOR3) == 1 {
		return 2
	} else if IoReadBit(SENSOR_FLOOR4) == 1 {
		return 3
	} else {
		return -1
	}
}

func Elev_get_stop_signal() bool {
	if Io_read_bit(STOP_BUTTON) == 1 {
		return true
	} else {
		return false
	}
}
func Elev_get_obstruction_signal() int {
	return Io_read_bit(OBSTRUCTION)
}
