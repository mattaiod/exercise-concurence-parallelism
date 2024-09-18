package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestFileSystem(t *testing.T) {
	fs := &FileSystem{files: make(map[string]string)}
	wg := sync.WaitGroup{}
	rand.Seed(time.Now().UnixNano())

	numWrites := rand.Intn(1000) + 1
	for i := 0; i < numWrites; i++ {
		filename := string(rand.Intn(26)+'a') + ".txt"
		content := string(rand.Intn(26) + 'a')
		wg.Add(1)
		go func(filename, content string) {
			defer wg.Done()
			fs.WriteFile(filename, content)
		}(filename, content)
	}

	wg.Wait()

	// Lire quelques fichiers pour vérifier leur contenu
	for i := 0; i < 10; i++ {
		filename := string(rand.Intn(26)+'a') + ".txt"
		fs.ReadFile(filename)
	}
}
