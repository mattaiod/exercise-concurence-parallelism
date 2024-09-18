package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestBank(t *testing.T) {
	bank := &Bank{}
	wg := sync.WaitGroup{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numDeposits := r.Intn(1000) + 1
	numWithdrawals := r.Intn(500) + 1 // Les retraits sont limités pour éviter un solde négatif

	totalDeposited := 0
	totalWithdrawn := 0

	// Effectuer des dépôts aléatoires
	for i := 0; i < numDeposits; i++ {
		amount := rand.Intn(1000) + 1
		totalDeposited += amount
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			bank.Deposit(amount)
		}(amount)
	}

	// Effectuer des retraits aléatoires
	for i := 0; i < numWithdrawals; i++ {
		amount := rand.Intn(1000) + 1
		totalWithdrawn += amount
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			bank.Withdraw(amount)
		}(amount)
	}

	wg.Wait()

	expectedBalance := totalDeposited - totalWithdrawn
	if bank.balance != expectedBalance {
		t.Errorf("Expected balance to be %d, but got %d", expectedBalance, bank.balance)
	}
}
