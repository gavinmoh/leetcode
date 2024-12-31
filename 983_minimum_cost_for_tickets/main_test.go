package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMincostTickets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		days := []int{1, 4, 6, 7, 8, 20}
		costs := []int{2, 7, 15}
		expected := 11

		assert.Equal(t, expected, mincostTickets(days, costs))
	})

	t.Run("test case 2", func(t *testing.T) {
		days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
		costs := []int{2, 7, 15}
		expected := 17

		assert.Equal(t, expected, mincostTickets(days, costs))
	})

	t.Run("test case 3", func(t *testing.T) {
		days := []int{1, 4, 6, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22, 23, 27, 28}
		costs := []int{3, 13, 45}
		expected := 44

		assert.Equal(t, expected, mincostTickets(days, costs))
	})
}
