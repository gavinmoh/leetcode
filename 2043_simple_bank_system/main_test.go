package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBank(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		balance := []int64{10, 100, 20, 50, 30}
		bank := Constructor(balance)

		assert.True(t, bank.Withdraw(3, 10))
		assert.True(t, bank.Transfer(5, 1, 20))
		assert.True(t, bank.Deposit(5, 20))
		assert.False(t, bank.Transfer(3, 4, 15))
		assert.False(t, bank.Withdraw(10, 50))
	})
}
