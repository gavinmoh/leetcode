package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCoupons(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		code := []string{"SAVE20", "", "PHARMA5", "SAVE@20"}
		businessLine := []string{"restaurant", "grocery", "pharmacy", "restaurant"}
		isActive := []bool{true, true, true, true}
		expected := []string{"PHARMA5", "SAVE20"}

		assert.Equal(t, expected, validateCoupons(code, businessLine, isActive))
	})

	t.Run("test case 2", func(t *testing.T) {
		code := []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"}
		businessLine := []string{"grocery", "electronics", "invalid"}
		isActive := []bool{false, true, true}
		expected := []string{"ELECTRONICS_50"}

		assert.Equal(t, expected, validateCoupons(code, businessLine, isActive))
	})
}
