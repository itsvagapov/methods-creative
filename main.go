package main

import "fmt"

func main() {
	bank := Bank{}

	id1 := bank.CreateAccount("Rasul", "USD", "debit", 10000)
	id2 := bank.CreateAccount("Muslim", "USD", "debit", 8000)
	
	bank.findAccount(1)
}
