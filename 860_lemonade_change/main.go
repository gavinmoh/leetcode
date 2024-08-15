package main

func lemonadeChange(bills []int) bool {
	cash := map[int]int{
		5:  0,
		10: 0,
		20: 0,
	}

	for _, bill := range bills {
		cash[bill] += 1
		change := bill - 5
		for change > 0 {
			if change == 5 {
				if cash[5] >= 1 {
					cash[5] -= 1
					change -= 5
					continue
				} else {
					return false
				}
			}

			if change == 10 {
				if cash[10] >= 1 {
					cash[10] -= 1
					change -= 10
					continue
				} else if cash[5] >= 1 {
					cash[5] -= 1
					change -= 5
					continue
				} else {
					return false
				}
			}

			if change == 15 {
				if cash[10] >= 1 {
					cash[10] -= 1
					change -= 10
					continue
				} else if cash[5] >= 1 {
					cash[5] -= 1
					change -= 5
				} else {
					return false
				}
			}
		}
	}

	return true
}
