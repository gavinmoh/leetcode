package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	case1 := []int{10, 15, 20}
	assert.Equal(t, 15, minCostClimbingStairs(case1))

	case2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	assert.Equal(t, 6, minCostClimbingStairs(case2))
}
