package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTime(t *testing.T) {
	case1Relations := [][]int{{1, 3}, {2, 3}}
	case1Time := []int{3, 2, 5}
	case1N := 3
	assert.Equal(t, 8, minimumTime(case1N, case1Relations, case1Time))

	case2Relations := [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}
	case2Time := []int{1, 2, 3, 4, 5}
	case2N := 5
	assert.Equal(t, 12, minimumTime(case2N, case2Relations, case2Time))

	case3Relations := [][]int{{2, 1}}
	case3Time := []int{1000, 1000}
	case3N := 2
	assert.Equal(t, 2000, minimumTime(case3N, case3Relations, case3Time))
}
