package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllPeople(t *testing.T) {
	t.Run("should return [0,1,2,3,5]", func(t *testing.T) {
		n := 6
		meetings := [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}
		firstPerson := 1
		expected := []int{0, 1, 2, 3, 5}

		assert.Equal(t, expected, findAllPeople(n, meetings, firstPerson))
	})

	t.Run("should return [0,1,3]", func(t *testing.T) {
		n := 4
		meetings := [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}
		firstPerson := 3
		expected := []int{0, 1, 3}

		assert.Equal(t, expected, findAllPeople(n, meetings, firstPerson))
	})

	t.Run("should return [0,1,2,3,4]", func(t *testing.T) {
		n := 5
		meetings := [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}
		firstPerson := 1
		expected := []int{0, 1, 2, 3, 4}

		assert.Equal(t, expected, findAllPeople(n, meetings, firstPerson))
	})
}
