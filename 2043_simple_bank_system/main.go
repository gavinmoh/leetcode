package main

type Bank struct {
	accounts []int64
	n        int
}

func Constructor(balance []int64) Bank {
	return Bank{
		accounts: append([]int64{0}, balance...),
		n:        len(balance),
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if this.ValidAccount(account1) && this.ValidAccount(account2) && this.accounts[account1] >= money {
		this.accounts[account1] -= money
		this.accounts[account2] += money

		return true
	}

	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if this.ValidAccount(account) {
		this.accounts[account] += money
		return true
	}

	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if this.ValidAccount(account) && this.accounts[account] >= money {
		this.accounts[account] -= money
		return true
	}

	return false
}

func (this *Bank) ValidAccount(account int) bool {
	return account <= this.n
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
