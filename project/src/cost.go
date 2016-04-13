//func costFunction() {

/*
5 heiser
4 etasjer

trykk utside etasje 2

input: knapp trykket, heisposisjoner for alle heiser
output: nærmeste tilgjengelige heis
elevator = costElevator (buttonchannel,)
*/

nmbr_of_elvtrs = 3

func costElevator (buttonchannel{type,floor} , lightChannel(a) , queSize(b)) {

	//button_matrix

	type State string

	select{
		switch {
		case IDLE:

			if (buttonchannel{up,0})
				case = 0_0
			else if (buttonchannel{up,1})
				case = 0_1
			else if (buttonchannel{up,2})
				case = 0_2
			else if (buttonchannel{down,1})
				case = 1_1
			else if (buttonchannel{down,2})
				case = 1_2
			else if (buttonchannel{down,3})
				case = 1_3
		

		case btn_0_0: // floor 0 opp
			min_distance = 3
			minQueSize = 100

			for i:=0 ; i<nmbr_of_elvtrs ; i++ {
				//Sjekk nærmeste heis til valgt etasje
				distance_to_elevator(i) = lightChannel(i) - floor(i) // heisposisjon - knappetasje ==> forskjell mellom heis og destinasjon
											// flere heiser: heisposisjon(i) - knappetasje ==> forskjell mellom heis og destinasjon
				if distance_to_elevator < min_distance {
					min_distance = distance_to_elevator
					closest_elevator = i
				}
			
			//Sjekk kø for alle heiser
				if queSize(i) < minQueSize {
					minQueSize := queSize(i)
				}


			}
			

			//if available_elvtr(closest_elevator) == FALSE

			Else
				Tilkall/aktiver hj
				hj to e1
				return closest_heis
			if (buttonchannel{up,0})
				case = 0_0
			else if (buttonchannel{up,1})
				case = 0_1
			else if (buttonchannel{up,2})
				case = 0_2
			else if (buttonchannel{down,1})
				case = 1_1
			else if (buttonchannel{down,2})
				case = 1_2
			else if (buttonchannel{down,3})
				case = 1_3
		

		case btn_0_1: // floor 1 opp
				Sjekk nærmeste heis til e2
				Hvis aktiv
					Sjekk nest nærmeste
				Else
					Tilkall/aktiver
			if (buttonchannel{up,0})
				case = 0_0
			else if (buttonchannel{up,1})
				case = 0_1
			else if (buttonchannel{up,2})
				case = 0_2
			else if (buttonchannel{down,1})
				case = 1_1
			else if (buttonchannel{down,2})
				case = 1_2
			else if (buttonchannel{down,3})
				case = 1_3
		


		case btn_0_2: // floor 2 opp
				Sjekk nærmeste heis til e2
				Hvis aktiv
					Sjekk nest nærmeste
				Else
					Tilkall/aktiver
			if (buttonchannel{up,0})
				case = 0_0
			else if (buttonchannel{up,1})
				case = 0_1
			else if (buttonchannel{up,2})
				case = 0_2
			else if (buttonchannel{down,1})
				case = 1_1
			else if (buttonchannel{down,2})
				case = 1_2
			else if (buttonchannel{down,3})
				case = 1_3
		

		case btn_1_1: // floor 1 ned
				Sjekk nærmeste heis til e3
				Hvis aktiv
					Sjekk nest nærmeste
				Else
					Tilkall/aktiver
			if (buttonchannel{up,0})
				case = 0_0
			else if (buttonchannel{up,1})
				case = 0_1
			else if (buttonchannel{up,2})
				case = 0_2
			else if (buttonchannel{down,1})
				case = 1_1
			else if (buttonchannel{down,2})
				case = 1_2
			else if (buttonchannel{down,3})
				case = 1_3
		

		case btn_1_2: // floor 2 ned
				Sjekk nærmeste heis til e3
				Hvis aktiv
					Sjekk nest nærmeste
				Else
					Tilkall/aktiver
			if (buttonchannel{up,0})
				case = 0_0
			else if (buttonchannel{up,1})
				case = 0_1
			else if (buttonchannel{up,2})
				case = 0_2
			else if (buttonchannel{down,1})
				case = 1_1
			else if (buttonchannel{down,2})
				case = 1_2
			else if (buttonchannel{down,3})
				case = 1_3
		

		case btn_1_3: // floor 3 ned
				Sjekk nærmeste heis til e4
				Hvis aktiv
					Sjekk nest nærmeste
				Else
					Tilkall/aktiver
			if (buttonchannel{up,0})
				case = 0_0
			else if (buttonchannel{up,1})
				case = 0_1
			else if (buttonchannel{up,2})
				case = 0_2
			else if (buttonchannel{down,1})
				case = 1_1
			else if (buttonchannel{down,2})
				case = 1_2
			else if (buttonchannel{down,3})
				case = 1_3
		

		} 	// End switch

	} 		// End select

}


}








