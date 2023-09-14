package main

// first starting for JFK
// find next tickets
// pick the one that has the lowest lexical order
// repeat until no more tickets
// return the itenerary

func findItinerary(tickets [][]string) []string {
	starting := "JFK"
	itinerary, ok := findItineraryRecursive(tickets, starting)

	if !ok {
		return []string{}
	}

	return itinerary
}

func findItineraryRecursive(tickets [][]string, starting string) ([]string, bool) {
	if len(tickets) == 0 {
		return []string{}, true
	} else if len(tickets) == 1 {
		if tickets[0][0] == starting {
			return tickets[0], true
		}
		return []string{}, false
	}

	nextTickets := [][]string{}
	for _, tickets := range tickets {
		if tickets[0] == starting {
			nextTickets = append(nextTickets, tickets)
		}
	}
	nextTickets = sortTickets(nextTickets)
	for _, nextTicket := range nextTickets {
		// remove ticket from tickets
		tickets = removeTicket(tickets, nextTicket)
		// find itinerary for next ticket
		itinerary, found := findItineraryRecursive(tickets, nextTicket[1])
		if found {
			return append([]string{starting}, itinerary...), true
		}
		// add ticket back to tickets
		tickets = append(tickets, nextTicket)
	}

	return []string{}, false
}

func sortTickets(tickets [][]string) [][]string {
	for i := 0; i < len(tickets); i++ {
		for j := i; j < len(tickets); j++ {
			if tickets[i][0] > tickets[j][0] {
				tickets[i], tickets[j] = tickets[j], tickets[i]
			} else if tickets[i][0] == tickets[j][0] {
				if tickets[i][1] > tickets[j][1] {
					tickets[i], tickets[j] = tickets[j], tickets[i]
				}
			}
		}
	}
	return tickets
}

func removeTicket(tickets [][]string, ticket []string) [][]string {
	for i, t := range tickets {
		if t[0] == ticket[0] && t[1] == ticket[1] {
			tickets = append(tickets[:i], tickets[i+1:]...)
			break
		}
	}
	return tickets
}
