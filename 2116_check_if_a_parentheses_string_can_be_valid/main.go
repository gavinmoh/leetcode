package main

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}

	open, unlocked := []int{}, []int{}
	for i, bracket := range s {
		if locked[i] == '0' {
			unlocked = append(unlocked, i)
		} else if bracket == '(' {
			open = append(open, i)
		} else if bracket == ')' {
			if len(open) > 0 {
				open = open[:len(open)-1]
			} else if len(unlocked) > 0 {
				unlocked = unlocked[:len(unlocked)-1]
			} else {
				return false
			}
		}
	}

	for len(open) > 0 && len(unlocked) > 0 && open[len(open)-1] < unlocked[len(unlocked)-1] {
		open = open[:len(open)-1]
		unlocked = unlocked[:len(unlocked)-1]
	}

	return len(open) == 0
}
