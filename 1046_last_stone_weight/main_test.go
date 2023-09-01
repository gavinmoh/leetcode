package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {

	case1 := []int{2, 7, 4, 1, 8, 1}
	assert.Equal(t, 1, LastStoneWeight(case1))

	case2 := []int{1, 3}
	assert.Equal(t, 2, LastStoneWeight(case2))

	case3 := []int{2, 2}
	assert.Equal(t, 0, LastStoneWeight(case3))

	case4 := []int{4, 3, 4, 3, 2}
	assert.Equal(t, 2, LastStoneWeight(case4))

}
