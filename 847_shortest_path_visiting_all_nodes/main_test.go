package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPathLength(t *testing.T) {
	case1 := [][]int{{1, 2, 3}, {0}, {0}, {0}}
	assert.Equal(t, 4, shortestPathLength(case1))

	case2 := [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}}
	assert.Equal(t, 4, shortestPathLength(case2))
}
