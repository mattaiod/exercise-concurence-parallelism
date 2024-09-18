package main

import (
	"sync"
)

type Journal struct {
	mu      sync.RWMutex
	entries []string
}

func (j *Journal) AddEntry(entry string) {
	// Verrouiller en écriture, ajouter l'entrée
}

func (j *Journal) GetEntries() []string {
	// Verrouiller en lecture, retourner toutes les entrées
	return nil
}
