package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := &Cache{store: make(map[string]string)}
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())

	numGoroutines := rand.Intn(1000) + 1
	for i := 0; i < numGoroutines; i++ {
		key := string(rand.Intn(26) + 'a')   // Générer une clé aléatoire
		value := string(rand.Intn(26) + 'a') // Générer une valeur aléatoire
		wg.Add(1)
		go func(key, value string) {
			defer wg.Done()
			cache.Set(key, value)
		}(key, value)
	}

	wg.Wait()

	// Valider que les lectures fonctionnent correctement
	for i := 0; i < numGoroutines; i++ {
		key := string(rand.Intn(26) + 'a')
		cache.Get(key)
	}
}
