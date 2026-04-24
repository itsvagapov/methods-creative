package models

import (
	// "fmt"

)

func (b *Bank) FindAccount(id int) (*Account, bool) {
	for i := range b.accounts {
		if b.accounts[i].ID == id {
			return &b.accounts[i], true
		}
	}
	return nil, false
}