package main

import (
	"sync"
)

type FileSystem struct {
	mu    sync.RWMutex
	files map[string]string
}

func (fs *FileSystem) WriteFile(filename, content string) {
	// Verrouiller en écriture, écrire dans le fichier
}

func (fs *FileSystem) ReadFile(filename string) string {
	// Verrouiller en lecture, retourner le contenu
	return ""
}
