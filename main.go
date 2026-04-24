package main

import (
	"fmt"

	"github.com/itsvagapov/methods-creative/models"
)

func main() {
	bank := models.Bank{}

	id1 := bank.CreateAccount("Muslim", "USD", "debit", 10000)
	id2 := bank.CreateAccount("Rasul", "USD", "debit", 8000)

	fmt.Println("=== Deposit ===")
	bank.Deposit(id1, 5000)
	bank.Deposit(id2, 3000)

	fmt.Println("=== Transfer ===")
	bank.Transfer(id1, id2, 1000)

	fmt.Println("=== Withdraw ===")
	bank.Withdraw(id2, 200)

	fmt.Println("=== Block account ===")
	bank.BlockAccount(id2)

	fmt.Println("=== Accounts ===")
	bank.ListAccounts()

	fmt.Println("=== Total balance ===")
	fmt.Println(bank.GetTotalBalance())
}