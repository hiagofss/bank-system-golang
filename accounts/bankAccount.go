package accounts

import "workspace/fiap/alura/go_course/src/bank/customers"

type BankAccount struct {
	Owner customers.Customer

	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *BankAccount) Withdraw(value float64) string {
	if value <= c.Balance && value > 0 {
		c.Balance -= value
		return "Withdrawal successfully completed"
	} else {
		return "Insufficient funds"
	}
}

func (c *BankAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.Balance += value
		return "Deposit successfully completed", c.Balance
	} else {
		return "Invalid deposit value", c.Balance
	}
}

func (c *BankAccount) Transfer(value float64, destiny *BankAccount) bool {
	if value <= c.Balance && value > 0 {
		c.Balance -= value
		destiny.Balance += value
		return true
	} else {
		return false
	}
}
