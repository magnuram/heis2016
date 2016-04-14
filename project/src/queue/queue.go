package queue

import (
	"driver"
)

var light driver.ElevLight
var elevInfo driver.ElevInfo

func QueueSetLights(lightChannel chan<- driver.ElevLight) {
	for i := 0; i < driver.N_FLOORS; i++ {
		lightChannel <- driver.ElevLight{Type: BUTTON_COMMAND, Floor: i, Active: elevInfo.ReqDown[i]}
		if i != 0 {
			lightChannel <- driver.ElevLight{Type: BUTTON_CALL_DOWN, Floor: i, Active: elevInfo.ReqUp[i]}
		}
		if i != 3 {
			lightChannel <- driver.ElevLight{Type: BUTTON_CALL_DOWN, Floor: i, Active: elevInfo.ReqUp[i]}
		}
	}
}
