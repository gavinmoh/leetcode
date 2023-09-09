package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum4(t *testing.T) {
	case1 := []int{1, 2, 3}
	assert.Equal(t, 7, CombinationSum4(case1, 4))

	case2 := []int{9}
	assert.Equal(t, 0, CombinationSum4(case2, 3))
}
