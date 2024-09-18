package main

import (
	"sync"
)

func updateMap(m map[int]int, key, value int, mu *sync.Mutex, wg *sync.WaitGroup) {
	// Verrouiller le mutex, ajouter la clé/valeur à la map, puis déverrouiller
}
