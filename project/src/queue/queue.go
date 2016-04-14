package queue

import (
	"driver"
)

//var light driver.ElevLight
var elevInfo driver.ElevInfo

func QueueSetLights(lightChannel <-chan driver.ElevLight) {
	var ligt driver.ElevLight
	ligt = <-lightChannel
	for i := 0; i < driver.N_FLOORS; i++ {
		lightChannel <- ligt{Type: BUTTON_COMMAND, Floor: i, Active: elevInfo.ReqDown[i]}
		if i != 0 {
			lightChannel <- ligt{Type: BUTTON_CALL_DOWN, Floor: i, Active: elevInfo.ReqUp[i]}
		}
		if i != 3 {
			lightChannel <- ligt{Type: BUTTON_CALL_DOWN, Floor: i, Active: elevInfo.ReqUp[i]}
		}
	}
}
