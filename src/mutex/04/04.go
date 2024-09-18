package main

import (
	"sync"
)

type ComplexStruct struct {
	mu   sync.Mutex
	data map[string]int
}

func (cs *ComplexStruct) Update(key string, value int) {
	// Verrouiller le mutex, modifier la structure, puis déverrouiller
}

func (cs *ComplexStruct) Get(key string) int {
	// Verrouiller en lecture, retourner la valeur, puis déverrouiller
	return 1
}
