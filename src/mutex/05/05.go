package main

import (
	"sync"
)

type Account struct {
	mu      sync.Mutex
	balance int
}

func (a *Account) Deposit(amount int) {
	// Verrouiller le mutex, déposer le montant
}

func (a *Account) Withdraw(amount int) bool {
	// Verrouiller le mutex, retirer le montant
	return false
}

func (a *Account) Transfer(to *Account, amount int) bool {
	// Utiliser des mutex pour sécuriser les transferts entre comptes
	return false
}
