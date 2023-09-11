package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupThePeople(t *testing.T) {
	case1 := []int{3, 3, 3, 3, 3, 1, 3}
	case1Expected := [][]int{{5}, {0, 1, 2}, {3, 4, 6}}
	case1Result := GroupThePeople(case1)
	for _, group := range case1Result {
		assert.Contains(t, case1Expected, group)
	}

	case2 := []int{2, 1, 3, 3, 3, 2}
	case2Expected := [][]int{{1}, {0, 5}, {2, 3, 4}}
	case2Result := GroupThePeople(case2)
	for _, group := range case2Result {
		assert.Contains(t, case2Expected, group)
	}
}
