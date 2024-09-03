package main

import "fmt"

func getLucky(s string, k int) int {
	str := ""
	for _, letter := range s {
		position := (letter - 'a') + 1
		str += fmt.Sprintf("%d", position)
	}

	var sum int
	for i := 0; i < k; i++ {
		sum = 0
		for _, letter := range str {
			sum += int(letter - '0')
		}
		str = fmt.Sprintf("%d", sum)
	}

	return sum
}
