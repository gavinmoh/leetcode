package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostConnectPoints(t *testing.T) {
	case1 := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	assert.Equal(t, 20, minCostConnectPoints(case1))

	case2 := [][]int{{3, 12}, {-2, 5}, {-4, 1}}
	assert.Equal(t, 18, minCostConnectPoints(case2))
}
