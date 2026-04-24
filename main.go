package main

import (
	"fmt"

	"github.com/itsvagapov/methods-creative/models"
)

func main() {
	bank := models.Bank{}

	// id1 := bank.CreateAccount("Rasul", "USD", "debit", 10000)
	// id2 := bank.CreateAccount("Muslim", "USD", "debit", 8000)

	fmt.Println(bank.FindAccount(2))
}
