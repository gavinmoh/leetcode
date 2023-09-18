package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKWeakestRows(t *testing.T) {
	case1 := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	assert.Equal(t, []int{2, 0, 3}, kWeakestRows(case1, 3))

	case2 := [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}
	assert.Equal(t, []int{0, 2}, kWeakestRows(case2, 2))
}
