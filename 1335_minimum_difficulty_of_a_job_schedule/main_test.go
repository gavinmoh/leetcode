package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDifficulty(t *testing.T) {
	t.Run("should return 7", func(t *testing.T) {
		jobDifficulty := []int{6, 5, 4, 3, 2, 1}
		d := 2
		expected := 7
		assert.Equal(t, expected, minDifficulty(jobDifficulty, d))
	})

	t.Run("should return -1", func(t *testing.T) {
		jobDifficulty := []int{4, 4, 4}
		d := 4
		expected := -1
		assert.Equal(t, expected, minDifficulty(jobDifficulty, d))
	})

	t.Run("should return 3", func(t *testing.T) {
		jobDifficulty := []int{1, 1, 1}
		d := 3
		expected := 3
		assert.Equal(t, expected, minDifficulty(jobDifficulty, d))
	})

}
