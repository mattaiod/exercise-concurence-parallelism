package main

import (
	"sync"
)

// Lit des valeurs d'un channel, les transforme, puis les renvoie dans un autre channel
func square(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	// Implémente le traitement de pipeline
}
