package main

import (
	"sync"
)

// Envoie des valeurs dans un channel à partir de plusieurs producteurs
func produce(ch chan int, value int, wg *sync.WaitGroup) {
	// Implémente les producteurs qui envoient des données dans le channel
}
