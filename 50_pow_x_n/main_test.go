package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyPow(t *testing.T) {
	case1x := 2.00000
	case1n := 10
	assert.Equal(t, math.Pow(case1x, float64(case1n)), myPow(case1x, case1n))

	case2x := 2.10000
	case2n := 3
	assert.Equal(t, math.Pow(case2x, float64(case2n)), myPow(case2x, case2n))

	case3x := 2.00000
	case3n := -2
	assert.Equal(t, math.Pow(case3x, float64(case3n)), myPow(case3x, case3n))

}
