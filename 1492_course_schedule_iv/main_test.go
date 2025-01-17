package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIfPrerequisite(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		numCourses := 2
		prerequisites := [][]int{{1, 0}}
		queries := [][]int{{0, 1}, {1, 0}}
		expected := []bool{false, true}

		assert.Equal(t, expected, checkIfPrerequisite(numCourses, prerequisites, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		numCourses := 2
		prerequisites := [][]int{}
		queries := [][]int{{1, 0}, {0, 1}}
		expected := []bool{false, false}

		assert.Equal(t, expected, checkIfPrerequisite(numCourses, prerequisites, queries))
	})

	t.Run("test case 3", func(t *testing.T) {
		numCourses := 3
		prerequisites := [][]int{{1, 2}, {1, 0}, {2, 0}}
		queries := [][]int{{1, 0}, {1, 2}}
		expected := []bool{true, true}

		assert.Equal(t, expected, checkIfPrerequisite(numCourses, prerequisites, queries))
	})
}
