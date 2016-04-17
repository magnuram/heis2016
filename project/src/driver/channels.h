#pragma once


//in port 4
#define Port4Subdevice        3
#define Port4ChannelOffset   16
#define Port4Direction        COMEDI_INPUT
#define OBSTRUCTION             (0x300+23)
#define STOP                    (0x300+22)
#define ButtonLocal1         (0x300+21)
#define ButtonLocal2         (0x300+20)
#define ButtonLocal3         (0x300+19)
#define ButtonLocal4         (0x300+18)
#define ButtonUp1              (0x300+17)
#define ButtonUp2              (0x300+16)

//in port 1
#define Port1Subdevice        2
#define Port1ChannelOffset   0
#define Port1Direction        COMEDI_INPUT
#define ButtonDown2            (0x200+0)
#define ButtonUp3              (0x200+1)
#define ButtonDown3            (0x200+2)
#define ButtonDown4            (0x200+3)
#define SensorFloor1           (0x200+4)
#define SensorFloor2           (0x200+5)
#define SensorFloor3           (0x200+6)
#define SensorFloor4           (0x200+7)

//out port 3
#define Port3Subdevice        3
#define Port3ChannelOffset   8
#define Port3Direction        COMEDI_OUTPUT
#define MOTORDIR                (0x300+15)
#define LightStop              (0x300+14)
#define LightCommand1          (0x300+13)
#define LightCommand2          (0x300+12)
#define LightCommand3          (0x300+11)
#define LightCommand4          (0x300+10)
#define LightUp1               (0x300+9)
#define LightUp2               (0x300+8)

//out port 2
#define Port2Subdevice        3
#define Port2ChannelOffset   0
#define Port2Direction        COMEDI_OUTPUT
#define LightDown2             (0x300+7)
#define LightUp3               (0x300+6)
#define LightDown3             (0x300+5)
#define LightDown4             (0x300+4)
#define LightDoorOpen         (0x300+3)
#define LightFloorInd2        (0x300+1)
#define LightFloorInd1        (0x300+0)

//out port 0
#define MOTOR                   (0x100+0)

//non-existing ports (for alignment)
#define ButtonDown1            -1
#define ButtonUp4              -1
#define LightDown1             -1
#define LightUp4               -1

