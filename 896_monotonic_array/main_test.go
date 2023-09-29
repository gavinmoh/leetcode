package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMonotonic(t *testing.T) {
	case1 := []int{1, 2, 2, 3}
	assert.True(t, isMonotonic(case1))

	case2 := []int{6, 5, 4, 4}
	assert.True(t, isMonotonic(case2))

	case3 := []int{1, 3, 2}
	assert.False(t, isMonotonic(case3))

	case4 := []int{1, 1, 0}
	assert.True(t, isMonotonic(case4))
}
