package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestAtomicIncrementWithLimit(t *testing.T) {
	var counter int64
	wg := sync.WaitGroup{}
	limit := int64(1000)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numGoroutines := r.Intn(2000) + 1

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go atomicIncrementWithLimit(&counter, limit, &wg)
	}

	wg.Wait()

	if counter != limit {
		t.Errorf("Expected counter to be %d, but got %d", limit, counter)
	}
}
