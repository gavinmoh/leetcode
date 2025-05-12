package main

import "sort"

func findEvenNumbers(digits []int) []int {
	n := len(digits)
	ints := map[int]bool{}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				a, b, c := digits[i], digits[j], digits[k]

				// a as first digit
				if a != 0 {
					ints[(a*100)+(b*10)+c] = true
					ints[(a*100)+(c*10)+b] = true
				}

				// b as first digit
				if b != 0 {
					ints[(b*100)+(a*10)+c] = true
					ints[(b*100)+(c*10)+a] = true
				}

				// c as first digit
				if c != 0 {
					ints[(c*100)+(a*10)+b] = true
					ints[(c*100)+(b*10)+a] = true
				}
			}
		}
	}

	result := []int{}
	for num, _ := range ints {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	sort.Ints(result)

	return result
}
