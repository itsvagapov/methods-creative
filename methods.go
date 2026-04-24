package main

import "fmt"


package main

type Account struct {
	ID         int
	Owner      string
	Balance    int
	Currency   string
	IsBlocked  bool
	DailyLimit int
	SpentToday int
	CreatedAt  string
	Type       string
	Status     string
}

type Bank struct {
	accounts []Account
	nextID   int
}