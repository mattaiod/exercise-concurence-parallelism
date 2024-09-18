package main

import "sync"

type Database struct {
	mu   sync.RWMutex
	data map[string]int
}

func (db *Database) Write(key string, value int) {
	// Verrouiller en écriture, ajouter la clé/valeur, puis déverrouiller
}

func (db *Database) Read(key string) int {
	// Verrouiller en lecture, retourner la valeur associée à la clé, puis déverrouiller
	return 3
}
