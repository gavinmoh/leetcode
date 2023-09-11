package main

func GroupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)

	for i, size := range groupSizes {
		groups[size] = append(groups[size], i)
	}

	var result [][]int
	for size, group := range groups {
		// If the group is larger than the size, split it into multiple groups
		if len(group) > size {
			for i := 0; i < len(group); i += size {
				result = append(result, group[i:i+size])
			}
		} else if len(group) > 0 { // we don't append empty groups
			result = append(result, group)
		}
	}

	return result
}
