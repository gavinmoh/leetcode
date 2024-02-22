package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindJudge(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		n := 2
		trust := [][]int{{1, 2}}
		expected := 2

		assert.Equal(t, expected, findJudge(n, trust))
	})

	t.Run("should return 3", func(t *testing.T) {
		n := 3
		trust := [][]int{{1, 3}, {2, 3}}
		expected := 3

		assert.Equal(t, expected, findJudge(n, trust))
	})

	t.Run("should return -1", func(t *testing.T) {
		n := 3
		trust := [][]int{{1, 3}, {2, 3}, {3, 1}}
		expected := -1

		assert.Equal(t, expected, findJudge(n, trust))
	})
}
