package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestChair(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		times := [][]int{{1, 4}, {2, 3}, {4, 6}}
		targetFriend := 1
		expected := 1

		assert.Equal(t, expected, smallestChair(times, targetFriend))
	})

	t.Run("test case 2", func(t *testing.T) {
		times := [][]int{{3, 10}, {1, 5}, {2, 6}}
		targetFriend := 0
		expected := 2

		assert.Equal(t, expected, smallestChair(times, targetFriend))
	})

	t.Run("test case 3", func(t *testing.T) {
		times := [][]int{
			{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310},
			{78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856},
			{98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821},
			{48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467},
		}
		targetFriend := 6
		expected := 2

		assert.Equal(t, expected, smallestChair(times, targetFriend))
	})
}
