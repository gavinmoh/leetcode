package main

func numberOfWays(corridor string) int {
	n := len(corridor)
	modulo := 1_000_000_007
	cache := make(map[int]map[int]int)

	var count func(index, seats int) int
	count = func(index, seats int) int {
		if index == n {
			if seats == 2 {
				return 1
			}
			return 0
		}

		if _, ok := cache[index]; ok {
			if _, ok := cache[index][seats]; ok {
				return cache[index][seats]
			}
		} else {
			newSeats := make(map[int]int)
			cache[index] = newSeats
		}

		result := 0
		if seats == 2 {
			if corridor[index] == 'S' {
				result = count(index+1, 1)
			} else {
				result = (count(index+1, 0) + count(index+1, 2)) % modulo
			}
		} else {
			if corridor[index] == 'S' {
				result = count(index+1, seats+1)
			} else {
				result = count(index+1, seats)
			}
		}

		cache[index][seats] = result
		return cache[index][seats]
	}

	return count(0, 0)

}
