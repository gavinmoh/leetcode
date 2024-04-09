package main

func timeRequiredToBuy(tickets []int, k int) int {
	count := 0
outer:
	for {
		for person := range tickets {
			if tickets[person] > 0 {
				tickets[person]--
				count++
			}
			if person == k && tickets[person] == 0 {
				break outer
			}
		}
	}

	return count
}
