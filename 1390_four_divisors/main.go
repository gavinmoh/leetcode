package main

import "math"

func sumFourDivisors(nums []int) int {
	cache := map[int]int{}
	result := 0

	for _, num := range nums {
		if _, ok := cache[num]; !ok {
			divisors := getDivisors(num)
			if len(divisors) == 4 {
				sum := 0
				for _, n := range divisors {
					sum += n
				}

				cache[num] = sum
			} else {
				cache[num] = 0
			}
		}

		result += cache[num]
	}

	return result
}

func getDivisors(n int) []int {
	divisors := map[int]struct{}{}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			divisors[i] = struct{}{}
			divisors[n/i] = struct{}{}
		}
	}

	result := []int{}
	for divisor := range divisors {
		result = append(result, divisor)
	}

	return result
}
