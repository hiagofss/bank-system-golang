package main

import (
	"fmt"
	"workspace/fiap/alura/go_course/src/bank/accounts"
)

func main() {
	acc1 := accounts.BankAccount{"Acc 1", 123, 456, 100.50}

	acc2 := accounts.BankAccount{Owner: "Acc 3", AgencyNumber: 123, AccountNumber: 456, Balance: 100.50}

	var acc3 *accounts.BankAccount
	acc3 = new(accounts.BankAccount)
	acc3.Owner = "Acc 3"
	acc3.AgencyNumber = 123
	acc3.AccountNumber = 456
	acc3.Balance = 100.50

	fmt.Println(acc1)
	fmt.Println(acc2)
	fmt.Println(acc3)

	acc3.Deposit(100)
	acc3.Transfer(200, &acc1)

	fmt.Println(acc1)
	fmt.Println(acc2)
	fmt.Println(acc3)
}
