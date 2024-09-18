package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestUpdateMap(t *testing.T) {
	var mu sync.Mutex
	sharedMap := make(map[int]int)
	wg := sync.WaitGroup{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numGoroutines := r.Intn(1000) + 1 // Générer un nombre aléatoire de goroutines

	expectedMap := make(map[int]int) // Crée une map pour stocker les valeurs attendues
	for i := 0; i < numGoroutines; i++ {
		key := rand.Intn(1000) // Générer une clé aléatoire
		value := rand.Intn(1000)
		expectedMap[key] = value

		wg.Add(1)
		go updateMap(sharedMap, key, value, &mu, &wg)
	}

	wg.Wait()

	// Comparer le contenu de sharedMap avec expectedMap
	for key, expectedValue := range expectedMap {
		if sharedMap[key] != expectedValue {
			t.Errorf("Expected sharedMap[%d] to be %d, but got %d", key, expectedValue, sharedMap[key])
		}
	}
}
