package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurvivedRobotsHealths(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		positions := []int{5, 4, 3, 2, 1}
		healths := []int{2, 17, 9, 15, 10}
		directions := "RRRRR"
		expected := []int{2, 17, 9, 15, 10}

		assert.Equal(t, expected, survivedRobotsHealths(positions, healths, directions))
	})

	t.Run("test case 2", func(t *testing.T) {
		positions := []int{3, 5, 2, 6}
		healths := []int{10, 10, 15, 12}
		directions := "RLRL"
		expected := []int{14}

		assert.Equal(t, expected, survivedRobotsHealths(positions, healths, directions))
	})

	t.Run("test case 3", func(t *testing.T) {
		positions := []int{1, 2, 5, 6}
		healths := []int{10, 10, 11, 11}
		directions := "RLRL"
		expected := []int{}

		assert.Equal(t, expected, survivedRobotsHealths(positions, healths, directions))
	})

	t.Run("test case 4", func(t *testing.T) {
		positions := []int{5, 46, 12}
		healths := []int{3, 27, 43}
		directions := "RLL"
		expected := []int{27, 42}

		assert.Equal(t, expected, survivedRobotsHealths(positions, healths, directions))
	})
}
