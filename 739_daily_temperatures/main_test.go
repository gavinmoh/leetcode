package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	t.Run("should return [1, 1, 4, 2, 1, 1, 0, 0]", func(t *testing.T) {
		temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
		expected := []int{1, 1, 4, 2, 1, 1, 0, 0}
		assert.Equal(t, expected, dailyTemperatures(temperatures))
	})

	t.Run("should return [1,1,1,0]", func(t *testing.T) {
		temperatures := []int{30, 40, 50, 60}
		expected := []int{1, 1, 1, 0}
		assert.Equal(t, expected, dailyTemperatures(temperatures))
	})

	t.Run("should return [1,1,0]", func(t *testing.T) {
		temperatures := []int{30, 60, 90}
		expected := []int{1, 1, 0}
		assert.Equal(t, expected, dailyTemperatures(temperatures))
	})

}
