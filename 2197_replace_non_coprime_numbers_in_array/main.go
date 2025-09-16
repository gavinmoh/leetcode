package main

func replaceNonCoprimes(nums []int) []int {
	processed := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		y := nums[i]
		for len(processed) > 0 {
			n := len(processed)
			x := processed[n-1]
			processed = processed[:n-1]

			gcd := findGCD(x, y)
			if gcd > 1 {
				y = lcm(x, y, gcd)
			} else {
				processed = append(processed, x)
				break
			}
		}

		processed = append(processed, y)
	}

	return processed
}

func findGCD(x, y int) int {
	if x == 0 {
		return y
	}

	return findGCD(y%x, x)
}

func lcm(x, y, gcd int) int {
	if x == 0 || y == 0 {
		return 0
	}
	return x * y / gcd
}
