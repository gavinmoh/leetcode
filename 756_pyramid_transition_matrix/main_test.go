package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPyramidTransition(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		bottom := "BCD"
		allowed := []string{"BCC", "CDE", "CEA", "FFF"}

		assert.True(t, pyramidTransition(bottom, allowed))
	})

	t.Run("test case 2", func(t *testing.T) {
		bottom := "AAAA"
		allowed := []string{"AAB", "AAC", "BCD", "BBE", "DEF"}

		assert.False(t, pyramidTransition(bottom, allowed))
	})
}
