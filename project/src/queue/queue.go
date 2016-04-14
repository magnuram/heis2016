package queue

import (
	//. "../driver"
	. "../config"
	//. "../elev"
)

func QOrderExist() bool {
	for i := 0; i < N_FLOORS; i++ {
		if Elinf.ReqLocal[i] != 0 || Elinf.ReqUp[i] != 0 || Elinf.ReqDown[i] != 0 {
			return true
		}
	}
	return false
}

func QSetLight(){
	for :=0 ; i < N_FLOORS; i++{
		ElevLight{Type:BUTTON_COMMAND ,Floor: i, Active: Elinf.ReqLocal}
		if i != 0{
		ElevLight{Type: BUTTON_CALL_DOWN,Floor: i, Active: Elinf.ReqDown}
		}
		if i != 3{
		ElevLight{Type: BUTTON_CALL_UP,Floor: i, Active: Elinf.ReqUp}
		}
	}


func QAddOrder(flr int, buttonchannel <-chan ElevButton) {

	select{
		case btn := <-buttonChannel
		switch btn.Type{
		case 0: //button call up 
			Elinf.ReqUp[flr] = 1
		case 1:
			Elinf.ReqDown[flr] = 1
		case 2:
			Elinf.ReqLocal[flr] = 1
		}
	}
}

func QOrdersAbove(currentFloor int) bool{
		for flr := currentFloor+1;  flr<N_FLOORS; flr++ {
		if Elinf.ReqLocal[i] != 0 || Elinf.ReqUp[i] != 0 || Elinf.ReqDown[i] != 0{
			return true
		} 
	}
	return false
}

func QOrdersBelow(currentFloor int) bool {
	for flr := currentFloor+1;  flr<currentFloor; flr++ {
		if Elinf.ReqLocal[i] != 0 || Elinf.ReqUp[i] != 0 || Elinf.ReqDown[i] != 0{
			return true
		} 
	}
	return false
	
}

func QChooseDir(currentFloor int, prevDir int)int{
	select{
		switch prevDir{
		case UP_DIR:
			if QOrdersAbove(currentFloor){
					return UP_DIR
				}else if QOrdersBelow(currentFloor){
					return DOWN_DIR
				}else{
					return STOP_DIR
				}

		case DOWN_DIR:
			if QOrdersAbove(currentFloor){
					return DOWN_DIR
				}else if QOrdersBelow(currentFloor){
					return UP_DIR
				}else{
					return STOP_DIR
				}

		case STOP_DIR:
			if QOrdersAbove(currentFloor){
					return UP_DIR
				}else if QOrdersBelow(currentFloor){
					return DOWN_DIR
				}else{
					return STOP_DIR
				}
		default:
			return STOP_DIR
		}
	}
}

func QShouldStop(flr int, prevDir int)int {
	if prevDir == -1{
		if Elinf.ReqDown[flr] != 0  || Elinf.ReqLocal[flr] != 0 || !QOrdersBelow(flr) || flr == 0{
			return 0
		}
	}
	if prevDir == 1{
		if Elinf.ReqUp[flr] != 0 || Elinf.ReqLocal[flr] != 0 || !QOrdersAbove(flr) || flr == 3{
			return 0
		}
	}
	return 1
}

func QDeleteOrders() {
	for i := 0; i < N_FLOORS; i++ {
		Elinf.ReqUP[i] = 0
		Elinf.ReqDown[i] = 0
		Elinf.ReqLocal[i] = 0
	}
	QSetLight()
	
}

func QDeleteManual(flr int){
		Elinf.ReqUP[flr] = 0
		Elinf.ReqDown[flr] = 0
		Elinf.ReqLocal[flr] = 0
		QSetLight()
}