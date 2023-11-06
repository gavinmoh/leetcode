package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeatManager(t *testing.T) {
	t.Run("should reserve seats", func(t *testing.T) {
		seatManager := Constructor(5)
		assert.Equal(t, 1, seatManager.Reserve())
		assert.Equal(t, 2, seatManager.Reserve())
		seatManager.Unreserve(2)
		assert.Equal(t, 2, seatManager.Reserve())
		assert.Equal(t, 3, seatManager.Reserve())
		assert.Equal(t, 4, seatManager.Reserve())
		assert.Equal(t, 5, seatManager.Reserve())
		seatManager.Unreserve(5)
	})
}
