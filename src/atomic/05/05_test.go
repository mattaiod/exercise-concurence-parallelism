package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestAtomicFlag(t *testing.T) {
	flag := &AtomicFlag{}
	wg := sync.WaitGroup{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numGoroutines := r.Intn(1000) + 1

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				flag.SetTrue()
			}
		}(i)
	}

	wg.Wait()

	if !flag.IsTrue() {
		t.Errorf("Expected flag to be true, but it is false")
	}
}
