package main

import (
	"sync"
)

// Crée un pool de workers pour traiter des tâches et envoyer des résultats
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	// Implémente la gestion de tâches concurrentes avec workers
}
