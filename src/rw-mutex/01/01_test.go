package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestReadMap(t *testing.T) {
	var mu sync.RWMutex
	sharedMap := make(map[int]int)
	rand.Seed(time.Now().UnixNano())

	// Remplir la map avec des valeurs aléatoires
	for i := 0; i < 1000; i++ {
		key := rand.Intn(1000)
		sharedMap[key] = i * 2
	}

	wg := sync.WaitGroup{}
	numReads := rand.Intn(1000) + 1

	for i := 0; i < numReads; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := rand.Intn(1000)
			val := readMap(sharedMap, key, &mu, &wg)
			expectedValue := sharedMap[key]
			if val != expectedValue {
				t.Errorf("Expected sharedMap[%d] to be %d, but got %d", key, expectedValue, val)
			}
		}(i)
	}

	wg.Wait()
}
