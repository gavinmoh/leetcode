package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMincostToHireWorkers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		quality := []int{10, 20, 5}
		wage := []int{70, 50, 30}
		k := 2
		expected := 105.00000

		assert.Equal(t, expected, mincostToHireWorkers(quality, wage, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		quality := []int{3, 1, 10, 10, 1}
		wage := []int{4, 8, 2, 2, 7}
		k := 3
		expected := 30.666666666666664

		assert.Equal(t, expected, mincostToHireWorkers(quality, wage, k))
	})
}
