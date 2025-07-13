package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchPlayersAndTrainers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		players := []int{4, 7, 9}
		trainers := []int{8, 2, 5, 8}
		expected := 2

		assert.Equal(t, expected, matchPlayersAndTrainers(players, trainers))
	})

	t.Run("test case 2", func(t *testing.T) {
		players := []int{1, 1, 1}
		trainers := []int{10}
		expected := 1

		assert.Equal(t, expected, matchPlayersAndTrainers(players, trainers))
	})
}
