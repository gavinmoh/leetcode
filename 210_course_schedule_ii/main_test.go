package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		numCourses := 2
		prerequisites := [][]int{{1, 0}}
		expected := []int{0, 1}

		assert.Equal(t, expected, findOrder(numCourses, prerequisites))
	})

	t.Run("test case 2", func(t *testing.T) {
		numCourses := 4
		prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
		expected := []int{0, 1, 2, 3}

		assert.Equal(t, expected, findOrder(numCourses, prerequisites))
	})

	t.Run("test case 3", func(t *testing.T) {
		numCourses := 1
		prerequisites := [][]int{}
		expected := []int{0}

		assert.Equal(t, expected, findOrder(numCourses, prerequisites))
	})
}
