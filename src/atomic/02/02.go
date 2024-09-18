package main

import (
	"sync"
)

// Définit de manière atomique une variable booléenne
func atomicSetBool(flag *int32, value int32, wg *sync.WaitGroup) {
	// Implémente le changement de valeur booléenne de manière atomique
}
