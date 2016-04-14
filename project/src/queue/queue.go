package queue

import (
	. "../driver"
)

//var light driver.ElevLight
var elevInfo driver.ElevInfo
var ligt driver.ElevLight

func QueueSetLights(lightChannel <-chan ElevLight) {

	ligt = <-lightChannel
	for i := 0; i < N_FLOORS; i++ {
		lightChannel <- ligt{Type: BUTTON_COMMAND, Floor: i, Active: elevInfo.ReqDown[i]}
		if i != 0 {
			lightChannel <- ligt{Type: BUTTON_CALL_DOWN, Floor: i, Active: elevInfo.ReqUp[i]}
		}
		if i != 3 {
			lightChannel <- ligt{Type: BUTTON_CALL_DOWN, Floor: i, Active: elevInfo.ReqUp[i]}
		}
	}
}
