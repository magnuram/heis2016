func costFunction() {
/*
5 heiser
4 etasjer

trykk utside etasje 2

input: knapp trykket, heisposisjoner for alle heiser
output: nærmeste tilgjengelige heis
elevator = costElevator (buttonchannel,)
*/

func costElevator (buttonchannel{type,floor}) {

	button_matrix

	type State string

	select{
		switch {
		case IDLE

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
		break

		case btn_0_0 // floor 0 opp
			min_distance = 3
			Sjekk nærmeste heis til e1
			for (i=1 ; i<=4 ; i++) {

				distance_to_heis = posisjon_heis(i) - floor

				if distance_to_heis < min
					min_distance = distance_to_heis(i)
					closest_heis = i
			}
			
			Hvis aktiv
			if available_elvtr(closest_heis) == FALSE

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
		break

		case btn_0_1 // floor 1 opp
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
		break


		case btn_0_2 // floor 2 opp
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
		break

		case btn_1_1 // floor 1 ned
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
		break

		case btn_1_2 // floor 2 ned
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
		break

		case btn_1_3 // floor 3 ned
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
		break

		} 	// End switch

	} 		// End select

}


}
