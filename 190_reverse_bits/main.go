package main

import (
	"strconv"
	"strings"
)

func reverseBits(n int) int {
	var bitStr strings.Builder

	for i := 0; i < 32; i++ {
		if n&1 == 1 {
			bitStr.WriteRune('1')
		} else {
			bitStr.WriteRune('0')
		}
		n >>= 1
	}

	result, _ := strconv.ParseInt(bitStr.String(), 2, 32)
	return int(result)
}
