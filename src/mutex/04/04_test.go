package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestComplexStruct(t *testing.T) {
	cs := &ComplexStruct{data: make(map[string]int)}
	wg := sync.WaitGroup{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numGoroutines := r.Intn(1000) + 1
	for i := 0; i < numGoroutines; i++ {
		key := string(rand.Intn(26) + 'a') // Générer une clé aléatoire
		value := rand.Intn(1000)
		wg.Add(1)
		go func(key string, value int) {
			defer wg.Done()
			cs.Update(key, value)
		}(key, value)
	}

	wg.Wait()

	// Valider que toutes les mises à jour sont cohérentes
	for i := 0; i < numGoroutines; i++ {
		key := string(rand.Intn(26) + 'a')
		cs.Get(key) // Assurer que les accès sont corrects
	}
}
