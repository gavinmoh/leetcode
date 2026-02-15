package main

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {
	aInt, _ := new(big.Int).SetString(a, 2)
	bInt, _ := new(big.Int).SetString(b, 2)
	sum := new(big.Int).Add(aInt, bInt)

	return fmt.Sprintf("%b", sum)
}
