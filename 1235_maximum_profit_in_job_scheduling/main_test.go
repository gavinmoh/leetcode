package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobScheduling(t *testing.T) {
	t.Run("should return 120", func(t *testing.T) {
		startTime := []int{1, 2, 3, 3}
		endTime := []int{3, 4, 5, 6}
		profit := []int{50, 10, 40, 70}
		expected := 120
		assert.Equal(t, expected, jobScheduling(startTime, endTime, profit))
	})

	t.Run("should return 150", func(t *testing.T) {
		startTime := []int{1, 2, 3, 4, 6}
		endTime := []int{3, 5, 10, 6, 9}
		profit := []int{20, 20, 100, 70, 60}
		expected := 150
		assert.Equal(t, expected, jobScheduling(startTime, endTime, profit))
	})
}
