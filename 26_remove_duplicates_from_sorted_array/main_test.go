package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {

	case1 := []int{1, 1, 2}
	case1Expected := []int{1, 2}
	assert.Equal(t, len(case1Expected), RemoveDuplicates(case1))
	for i, el := range case1Expected {
		assert.Equal(t, el, case1[i])
	}

	case2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	case2Expected := []int{0, 1, 2, 3, 4}
	assert.Equal(t, len(case2Expected), RemoveDuplicates(case2))
	for i, el := range case2Expected {
		assert.Equal(t, el, case2[i])
	}

}
