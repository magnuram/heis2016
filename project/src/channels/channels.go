// channel.go gives readable names to the subdevices/channels used with the elevator
package elev

const (
	//in port 4
	PORT4           = 3
	OBSTRUCTION     = (0x300 + 23)
	StopButton     = (0x300 + 22)
	ButtonLocal1 = (0x300 + 21)
	ButtonLocal2 = (0x300 + 20)
	ButtonLocal3 = (0x300 + 19)
	ButtonLocal4 = (0x300 + 18)
	ButtonUp1      = (0x300 + 17)
	ButtonUp2      = (0x300 + 16)

	//in port 1
	PORT1         = 2
	ButtonDown2  = (0x200 + 0)
	ButtonUp3    = (0x200 + 1)
	ButtonDown3  = (0x200 + 2)
	ButtonDown4  = (0x200 + 3)
	SensorFloor1 = (0x200 + 4)
	SensorFloor2 = (0x200 + 5)
	SensorFloor3 = (0x200 + 6)
	SensorFloor4 = (0x200 + 7)

	//out port 3
	PORT3          = 3
	MOTORDIR       = (0x300 + 15)
	LightStop     = (0x300 + 14)
	LightCommand1 = (0x300 + 13)
	LightCommand2 = (0x300 + 12)
	LightCommand3 = (0x300 + 11)
	LightCommand4 = (0x300 + 10)
	LightUp1      = (0x300 + 9)
	LightUp2      = (0x300 + 8)

	//out port 2
	PORT2            = 3
	LightDown2      = (0x300 + 7)
	LightUp3        = (0x300 + 6)
	LightDown3      = (0x300 + 5)
	LightDown4      = (0x300 + 4)
	LightDoorOpen  = (0x300 + 3)
	LightFloorInd2 = (0x300 + 1)
	LightFloorInd1 = (0x300 + 0)

	//out port 0
	PORT0 = 1
	MOTOR = (0x100 + 0)

	//non-existing ports (for alignment)
	ButtonDown1 = -1
	ButtonUp4   = -1
	LightDown1  = -1
	LightUp4    = -1
)
