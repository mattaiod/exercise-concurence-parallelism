package main

import (
	"sync"
)

func readMap(m map[int]int, key int, mu *sync.RWMutex, wg *sync.WaitGroup) int {
	// Verrouiller en lecture, retourner la valeur, puis déverrouiller
	return 3
}
