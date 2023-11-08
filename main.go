package main

import (
	"fmt"
	"workspace/fiap/alura/go_course/src/bank/accounts"
	"workspace/fiap/alura/go_course/src/bank/customers"
)

func PayBill(account verifyAccount, value float64) {
	account.Withdraw(value)
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {
	// acc1 := accounts.BankAccount{"Acc 1", 123, 456, 100.50}

	acc2 := accounts.BankAccount{Owner: customers.Customer{Name: "Acc 3", Document: "12345678910", JobTitle: "Developer"}, AgencyNumber: 123, AccountNumber: 456}

	customer3 := customers.Customer{Name: "Acc 3", Document: "12345678910", JobTitle: "Developer"}
	acc3 := accounts.SavingAccount{
		Owner:         customer3,
		AgencyNumber:  123,
		AccountNumber: 456,
		Operation:     1,
	}

	acc3.Owner = customer3
	acc3.AgencyNumber = 123
	acc3.AccountNumber = 456

	// fmt.Println(acc1)
	fmt.Println("acc2> ", acc2.GetBalance())
	fmt.Println("acc3> ", acc3.GetBalance())

	acc3.Deposit(100)
	// acc3.Transfer(200, &acc1)

	// fmt.Pri"acc1> ",ntln(acc1.GetBalance())
	fmt.Println("acc2> ", acc2.GetBalance())
	fmt.Println("acc3> ", acc3.GetBalance())

	PayBill(&acc3, 60)
	fmt.Println("acc3> ", acc3.GetBalance())

}
