package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	case1 := 2
	case1Expected := []int{0, 1, 1}
	assert.Equal(t, case1Expected, CountBits(case1))

	case2 := 5
	case2Expected := []int{0, 1, 1, 2, 1, 2}
	assert.Equal(t, case2Expected, CountBits(case2))
}
