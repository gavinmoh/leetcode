package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartitionLabels(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "ababcbacadefegdehijhklij"
		expected := []int{9, 7, 8}

		assert.Equal(t, expected, partitionLabels(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "eccbbbbdec"
		expected := []int{10}

		assert.Equal(t, expected, partitionLabels(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "caedbdedda"
		expected := []int{1, 9}

		assert.Equal(t, expected, partitionLabels(s))
	})
}
