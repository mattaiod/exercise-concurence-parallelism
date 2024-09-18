package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestJournal(t *testing.T) {
	journal := &Journal{}
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())

	numEntries := rand.Intn(1000) + 1
	for i := 0; i < numEntries; i++ {
		entry := string(rand.Intn(26) + 'a') // Générer une entrée aléatoire
		wg.Add(1)
		go func(entry string) {
			defer wg.Done()
			journal.AddEntry(entry)
		}(entry)
	}

	wg.Wait()

	entries := journal.GetEntries()
	if len(entries) != numEntries {
		t.Errorf("Expected %d entries in journal, but got %d", numEntries, len(entries))
	}
}
