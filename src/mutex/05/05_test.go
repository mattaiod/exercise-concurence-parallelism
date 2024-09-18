package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestAccountTransfer(t *testing.T) {
	acc1 := &Account{balance: 10000}
	acc2 := &Account{balance: 10000}
	wg := sync.WaitGroup{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numTransfers := r.Intn(1000) + 1

	for i := 0; i < numTransfers; i++ {
		amount := rand.Intn(500) + 1
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			if rand.Intn(2) == 0 {
				acc1.Transfer(acc2, amount)
			} else {
				acc2.Transfer(acc1, amount)
			}
		}(amount)
	}

	wg.Wait()

	// Vérifier que la somme totale des comptes est toujours correcte
	if acc1.balance+acc2.balance != 20000 {
		t.Errorf("Expected total balance to be 20000, but got %d", acc1.balance+acc2.balance)
	}
}
