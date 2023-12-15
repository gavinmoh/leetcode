package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDestCity(t *testing.T) {
	t.Run("should return Sao Paulo", func(t *testing.T) {
		paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
		expected := "Sao Paulo"
		assert.Equal(t, expected, destCity(paths))
	})

	t.Run("should return A", func(t *testing.T) {
		paths := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
		expected := "A"
		assert.Equal(t, expected, destCity(paths))
	})

	t.Run("should return Z", func(t *testing.T) {
		paths := [][]string{{"A", "Z"}}
		expected := "Z"
		assert.Equal(t, expected, destCity(paths))
	})

}
