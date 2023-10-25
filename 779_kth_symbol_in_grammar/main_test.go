package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthGrammar(t *testing.T) {
	t.Run("should return 0", func(t *testing.T) {
		assert.Equal(t, 0, kthGrammar(1, 1))
	})

	t.Run("should return 0", func(t *testing.T) {
		assert.Equal(t, 0, kthGrammar(2, 1))
	})

	t.Run("should return 1", func(t *testing.T) {
		assert.Equal(t, 1, kthGrammar(2, 2))
	})
}
