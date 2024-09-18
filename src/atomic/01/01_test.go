package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestAtomicIncrementCounter(t *testing.T) {
	var counter int64
	wg := sync.WaitGroup{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numGoroutines := r.Intn(1000) + 1 // Génère un nombre aléatoire de goroutines (entre 1 et 1000)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go atomicIncrementCounter(&counter, &wg)
	}

	wg.Wait()

	if counter != int64(numGoroutines) {
		t.Errorf("Expected counter to be %d, but got %d", numGoroutines, counter)
	}
}
