package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	t.Run("should return [[bat],[nat,tan],[ate,eat,tea]]", func(t *testing.T) {
		strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		expected := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}

		assert.Equal(t, expected, groupAnagrams(strs))
	})

	t.Run("should return [[]]", func(t *testing.T) {
		strs := []string{""}
		expected := [][]string{{""}}

		assert.Equal(t, expected, groupAnagrams(strs))
	})

	t.Run("should return [[a]]", func(t *testing.T) {
		strs := []string{"a"}
		expected := [][]string{{"a"}}

		assert.Equal(t, expected, groupAnagrams(strs))
	})
}
