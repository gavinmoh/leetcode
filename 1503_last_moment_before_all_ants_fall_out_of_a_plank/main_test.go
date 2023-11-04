package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLastMoment(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		n := 4
		left := []int{4, 3}
		right := []int{0, 1}
		expected := 4
		assert.Equal(t, expected, getLastMoment(n, left, right))
	})

	t.Run("should return 7", func(t *testing.T) {
		n := 7
		left := []int{}
		right := []int{0, 1, 2, 3, 4, 5, 6, 7}
		expected := 7
		assert.Equal(t, expected, getLastMoment(n, left, right))
	})

	t.Run("should return 7", func(t *testing.T) {
		n := 7
		left := []int{0, 1, 2, 3, 4, 5, 6, 7}
		right := []int{}
		expected := 7
		assert.Equal(t, expected, getLastMoment(n, left, right))
	})

	t.Run("should return 17", func(t *testing.T) {
		n := 20
		left := []int{4, 7, 15}
		right := []int{9, 3, 13, 10}
		expected := 17
		assert.Equal(t, expected, getLastMoment(n, left, right))
	})
}
