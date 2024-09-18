package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	tasks := make(chan int, 10)
	results := make(chan int, 10)
	wg := sync.WaitGroup{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numWorkers := r.Intn(10) + 1
	numTasks := r.Intn(100) + 1
	expectedSum := 0

	// Lancer des workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, tasks, results, &wg)
	}

	// Envoyer des tâches
	for i := 0; i < numTasks; i++ {
		tasks <- i
		expectedSum += i * 2
	}
	close(tasks)

	// Fermer le canal des résultats quand tous les workers sont terminés
	go func() {
		wg.Wait()
		close(results)
	}()

	// Calculer la somme des résultats
	sum := 0
	for result := range results {
		sum += result
	}

	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}
}
