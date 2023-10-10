package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	case1 := []int{4, 2, 5, 3}
	assert.Equal(t, 0, minOperations(case1))

	case2 := []int{1, 2, 3, 5, 6}
	assert.Equal(t, 1, minOperations(case2))

	case3 := []int{1, 10, 100, 1000}
	assert.Equal(t, 3, minOperations(case3))

	case4 := []int{8, 10, 16, 18, 10, 10, 16, 13, 13, 16}
	assert.Equal(t, 6, minOperations(case4))
}
