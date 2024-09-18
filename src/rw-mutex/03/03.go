package main

import "sync"

type Cache struct {
	mu    sync.RWMutex
	store map[string]string
}

func (c *Cache) Set(key, value string) {
	// Verrouiller en écriture, ajouter la clé/valeur
}

func (c *Cache) Get(key string) string {
	// Verrouiller en lecture, retourner la valeur
	return ""
}
