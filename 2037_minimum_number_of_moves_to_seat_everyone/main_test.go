package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMovesToSeat(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		seats := []int{3, 1, 5}
		students := []int{2, 7, 4}
		expected := 4

		assert.Equal(t, expected, minMovesToSeat(seats, students))
	})

	t.Run("test case 2", func(t *testing.T) {
		seats := []int{4, 1, 5, 9}
		students := []int{1, 3, 2, 6}
		expected := 7

		assert.Equal(t, expected, minMovesToSeat(seats, students))
	})

	t.Run("test case 3", func(t *testing.T) {
		seats := []int{2, 2, 6, 6}
		students := []int{1, 3, 2, 6}
		expected := 4

		assert.Equal(t, expected, minMovesToSeat(seats, students))
	})
}
