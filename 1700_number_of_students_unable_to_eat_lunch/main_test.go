package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountStudents(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		students := []int{1, 1, 0, 0}
		sandwiches := []int{0, 1, 0, 1}
		expected := 0

		assert.Equal(t, expected, countStudents(students, sandwiches))
	})

	t.Run("test case 2", func(t *testing.T) {
		students := []int{1, 1, 1, 0, 0, 1}
		sandwiches := []int{1, 0, 0, 0, 1, 1}
		expected := 3

		assert.Equal(t, expected, countStudents(students, sandwiches))
	})
}
