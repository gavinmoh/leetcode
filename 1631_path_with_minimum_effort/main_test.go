package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumEffortPath(t *testing.T) {
	case1 := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	assert.Equal(t, 2, minimumEffortPath(case1))

	case2 := [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	assert.Equal(t, 1, minimumEffortPath(case2))

	case3 := [][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}
	assert.Equal(t, 0, minimumEffortPath(case3))
}
