package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestEquivalentString(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s1 := "parker"
		s2 := "morris"
		baseStr := "parser"
		expected := "makkek"

		assert.Equal(t, expected, smallestEquivalentString(s1, s2, baseStr))
	})

	t.Run("test case 2", func(t *testing.T) {
		s1 := "hello"
		s2 := "world"
		baseStr := "hold"
		expected := "hdld"

		assert.Equal(t, expected, smallestEquivalentString(s1, s2, baseStr))
	})

	t.Run("test case 3", func(t *testing.T) {
		s1 := "leetcode"
		s2 := "programs"
		baseStr := "sourcecode"
		expected := "aauaaaaada"

		assert.Equal(t, expected, smallestEquivalentString(s1, s2, baseStr))
	})

	t.Run("test case 4", func(t *testing.T) {
		s1 := "dfeffdfafbbebbebacbbdfcfdbcacdcbeeffdfebbdebbdafff"
		s2 := "adcdfabadbeeafeabbadcefcaabdecabfecffbabbfcdfcaaae"
		baseStr := "myickvflcpfyqievitqtwvfpsrxigauvlqdtqhpfugguwfcpqv"
		expected := "myiakvalapayqiavitqtwvapsrxigauvlqatqhpaugguwaapqv"

		assert.Equal(t, expected, smallestEquivalentString(s1, s2, baseStr))
	})
}
