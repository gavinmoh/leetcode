package main

func garbageCollection(garbage []string, travel []int) int {
	paper, metal, glass := 0, 0, 0
	collectTime := 0

	for i, garb := range garbage {
		for _, g := range garb {
			collectTime += 1
			switch g {
			case 'P':
				if paper < i {
					paper = i
				}
			case 'M':
				if metal < i {
					metal = i
				}
			case 'G':
				if glass < i {
					glass = i
				}
			}
		}
	}

	travelTime := 0
	for i, time := range travel {
		if paper-1 >= i {
			travelTime += time
		}
		if metal-1 >= i {
			travelTime += time
		}
		if glass-1 >= i {
			travelTime += time
		}
	}

	return collectTime + travelTime

}
