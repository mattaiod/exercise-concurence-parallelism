package main

import (
	"sync"
	"testing"
)

func TestDatabase(t *testing.T) {
	db := &Database{data: make(map[string]int)}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				db.Write("key", i)
			}(i)
		} else {
			wg.Add(1)
			go func() {
				defer wg.Done()
				_ = db.Read("key")
			}()
		}
	}

	wg.Wait()

	if val := db.Read("key"); val == 0 {
		t.Errorf("Expected value in database to be non-zero")
	}
}
