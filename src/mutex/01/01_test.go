package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// TestIncrementCounter vérifie que la fonction incrementCounter incrémente correctement un compteur partagé entre plusieurs goroutines.
func TestIncrementCounter(t *testing.T) {
	var mu sync.Mutex
	var counter int
	wg := sync.WaitGroup{}

	// Crée une nouvelle source de randomisation et un nouvel objet rand
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Génère un nombre aléatoire de goroutines (entre 1 et 1000)
	numGoroutines := r.Intn(1000) + 1

	// Démarre les goroutines pour incrémenter le compteur
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementCounter(&counter, &mu, &wg)
	}

	// Attends que toutes les goroutines aient terminé
	wg.Wait()

	// Vérifie que le compteur est égal au nombre de goroutines
	if counter != numGoroutines {
		t.Errorf("Expected counter to be %d, but got %d", numGoroutines, counter)
	}
}
