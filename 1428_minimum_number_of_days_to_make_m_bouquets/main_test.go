package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		bloomDay := []int{1, 10, 3, 10, 2}
		m := 3
		k := 1
		expected := 3

		assert.Equal(t, expected, minDays(bloomDay, m, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		bloomDay := []int{1, 10, 3, 10, 2}
		m := 3
		k := 2
		expected := -1

		assert.Equal(t, expected, minDays(bloomDay, m, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		bloomDay := []int{7, 7, 7, 7, 12, 7, 7}
		m := 2
		k := 3
		expected := 12

		assert.Equal(t, expected, minDays(bloomDay, m, k))
	})
}
