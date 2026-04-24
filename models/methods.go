package models

import "fmt"

func (b *Bank) findAccount(id int) (*Account, bool) {
	for i := range b.accounts {
		if b.accounts[i].ID == id {
			return &b.accounts[i], true
		}
	}
	return nil, false
}

func (b *Bank) CreateAccount(owner, currency, accType string, limit int) int {
	acc := Account{
		ID:         b.nextID,
		Owner:      owner,
		Balance:    0,
		Currency:   currency,
		IsBlocked:  false,
		DailyLimit: limit,
		SpentToday: 0,
		Type:       accType,
		Status:     "active",
		CreatedAt:  "now",
	}

	b.accounts = append(b.accounts, acc)
	b.nextID++

	return acc.ID
}

func (b *Bank) Deposit(id, amount int) bool {
	if amount <= 0 {
		return false
	}

	acc, ok := b.findAccount(id)
	if !ok || acc.IsBlocked {
		return false
	}

	acc.Balance += amount
	acc.SpentToday += amount
	return true
}

func (b *Bank) Withdraw(id, amount int) bool {
	if amount <= 0 {
		return false
	}

	acc, ok := b.findAccount(id)
	if !ok || acc.IsBlocked {
		return false
	}

	if acc.Balance < amount {
		return false
	}

	if acc.SpentToday+amount > acc.DailyLimit {
		return false
	}

	acc.Balance -= amount
	acc.SpentToday += amount

	return true
}

func (b *Bank) Transfer(fromID, toID, amount int) bool {
	from, ok1 := b.findAccount(fromID)
	to, ok2 := b.findAccount(toID)

	if !ok1 || !ok2 {
		return false
	}

	if from.IsBlocked || to.IsBlocked {
		return false
	}

	fee := amount / 100
	total := amount + fee

	if from.Balance < total {
		return false
	}

	if from.SpentToday+total > from.DailyLimit {
		return false
	}

	from.Balance -= total
	to.Balance += amount
	from.SpentToday += total

	return true
}

func (b *Bank) BlockAccount(id int) bool {
	acc, ok := b.findAccount(id)
	if !ok {
		return false
	}

	acc.IsBlocked = true
	acc.Status = "blocked"
	return true
}

func (b *Bank) UnblockAccount(id int) bool {
	acc, ok := b.findAccount(id)
	if !ok {
		return false
	}

	acc.IsBlocked = false
	acc.Status = "active"
	return true
}

func (b *Bank) GetTotalBalance() int {
	total := 0
	for _, acc := range b.accounts {
		if !acc.IsBlocked {
			total += acc.Balance
		}
	}
	return total
}

func (b *Bank) ListAccounts() {
	for _, a := range b.accounts {
		fmt.Printf("ID:%d | %s | %d %s | blocked:%v | limit:%d\n",
			a.ID, a.Owner, a.Balance, a.Currency, a.IsBlocked, a.DailyLimit)
	}
}