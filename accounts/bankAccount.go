package accounts

import "workspace/fiap/alura/go_course/src/bank/customers"

type BankAccount struct {
	Owner         customers.Customer
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *BankAccount) Withdraw(value float64) string {
	if value <= c.balance && value > 0 {
		c.balance -= value
		return "Withdrawal successfully completed"
	} else {
		return "Insufficient funds"
	}
}

func (c *BankAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposit successfully completed", c.balance
	} else {
		return "Invalid deposit value", c.balance
	}
}

func (c *BankAccount) Transfer(value float64, destiny *BankAccount) bool {
	if value <= c.balance && value > 0 {
		c.balance -= value
		destiny.balance += value
		return true
	} else {
		return false
	}
}

func (c *BankAccount) GetBalance() float64 {
	return c.balance
}
