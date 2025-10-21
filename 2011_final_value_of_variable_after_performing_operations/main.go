package main

import "strings"

func finalValueAfterOperations(operations []string) int {
	x := 0

	for _, ops := range operations {
		if strings.Contains(ops, "+") {
			x++
		} else {
			x--
		}
	}

	return x
}
