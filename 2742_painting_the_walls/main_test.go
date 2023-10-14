package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaintWalls(t *testing.T) {
	case1Cost := []int{1, 2, 3, 2}
	case1Time := []int{1, 2, 3, 2}
	assert.Equal(t, 3, paintWalls(case1Cost, case1Time))

	case2Cost := []int{2, 3, 4, 2}
	case2Time := []int{1, 1, 1, 1}
	assert.Equal(t, 4, paintWalls(case2Cost, case2Time))
}
