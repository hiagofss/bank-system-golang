package accounts

import "workspace/fiap/alura/go_course/src/bank/customers"

type SavingAccount struct {
	Owner                                  customers.Customer
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingAccount) Withdraw(value float64) string {
	if value <= c.balance && value > 0 {
		c.balance -= value
		return "Withdrawal successfully completed"
	} else {
		return "Insufficient funds"
	}
}

func (c *SavingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposit successfully completed", c.balance
	} else {
		return "Invalid deposit value", c.balance
	}
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
