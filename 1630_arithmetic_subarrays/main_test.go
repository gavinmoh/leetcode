package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckArithmeticSubarrays(t *testing.T) {
	t.Run("should return [true,false,true]", func(t *testing.T) {
		nums := []int{4, 6, 5, 9, 3, 7}
		l := []int{0, 0, 2}
		r := []int{2, 3, 5}
		expected := []bool{true, false, true}
		assert.Equal(t, expected, checkArithmeticSubarrays(nums, l, r))
	})

	t.Run("should return [false,true,false,true,true,false]", func(t *testing.T) {
		nums := []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}
		l := []int{0, 1, 6, 4, 8, 7}
		r := []int{4, 4, 9, 7, 9, 10}
		expected := []bool{false, true, false, false, true, true}
		assert.Equal(t, expected, checkArithmeticSubarrays(nums, l, r))
	})
}
