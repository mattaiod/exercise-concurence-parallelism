package main

import "sync"

type Bank struct {
	mu      sync.Mutex
	balance int
}

func (b *Bank) Deposit(amount int) {
	// Verrouiller le mutex, déposer le montant, puis déverrouiller
}

func (b *Bank) Withdraw(amount int) bool {
	// Verrouiller le mutex, vérifier et retirer le montant, puis déverrouiller
	// Retourner false si le solde est insuffisant
	return false
}
