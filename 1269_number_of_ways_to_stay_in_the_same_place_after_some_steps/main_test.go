package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumWays(t *testing.T) {
	case1Steps := 3
	case1ArrLen := 2
	assert.Equal(t, 4, numWays(case1Steps, case1ArrLen))

	case2Steps := 2
	case2ArrLen := 4
	assert.Equal(t, 2, numWays(case2Steps, case2ArrLen))

	case3Steps := 4
	case3ArrLen := 2
	assert.Equal(t, 8, numWays(case3Steps, case3ArrLen))
}
