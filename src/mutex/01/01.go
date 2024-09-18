package main

import (
	"sync"
)

func incrementCounter(counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	// Verrouiller le mutex, incrémenter le compteur, puis déverrouiller
}
