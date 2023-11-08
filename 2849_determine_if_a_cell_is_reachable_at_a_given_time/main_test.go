package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsReachableAtTime(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		sx, sy := 2, 4
		fx, fy := 7, 7
		time := 6
		assert.True(t, isReachableAtTime(sx, sy, fx, fy, time))
	})

	t.Run("should return false", func(t *testing.T) {
		sx, sy := 3, 1
		fx, fy := 7, 3
		time := 3
		assert.False(t, isReachableAtTime(sx, sy, fx, fy, time))
	})
}
