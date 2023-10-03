package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIdenticalPairs(t *testing.T) {
	case1 := []int{1, 2, 3, 1, 1, 3}
	assert.Equal(t, 4, numIdenticalPairs(case1))

	case2 := []int{1, 1, 1, 1}
	assert.Equal(t, 6, numIdenticalPairs(case2))

	case3 := []int{1, 2, 3}
	assert.Equal(t, 0, numIdenticalPairs(case3))
}
