package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageSmoother(t *testing.T) {
	t.Run("should return [[0,0,0],[0,0,0],[0,0,0]]", func(t *testing.T) {
		img := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
		expected := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
		assert.Equal(t, expected, imageSmoother(img))
	})

	t.Run("should return [[137,141,137],[141,138,141],[137,141,137]]", func(t *testing.T) {
		img := [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}
		expected := [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}}
		assert.Equal(t, expected, imageSmoother(img))
	})
}
