package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoorPigs(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		assert.Equal(t, 2, poorPigs(4, 15, 15))
		assert.Equal(t, 2, poorPigs(4, 15, 30))
	})

	t.Run("should return 5", func(t *testing.T) {
		assert.Equal(t, 5, poorPigs(1000, 15, 60))
	})

	t.Run("should return 3", func(t *testing.T) {
		assert.Equal(t, 3, poorPigs(125, 1, 4))
	})
}
