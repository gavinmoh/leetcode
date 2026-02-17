package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadBinaryWatch(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		turnedOn := 1
		expected := []string{"0:32", "0:16", "0:08", "0:04", "0:02", "0:01", "8:00", "4:00", "2:00", "1:00"}

		assert.Equal(t, expected, readBinaryWatch(turnedOn))
	})

	t.Run("test case 2", func(t *testing.T) {
		turnedOn := 9
		expected := []string{}

		assert.Equal(t, expected, readBinaryWatch(turnedOn))
	})
}
