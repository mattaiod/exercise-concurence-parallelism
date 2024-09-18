package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomicCompareAndSwap(t *testing.T) {
	var lock int32
	wg := sync.WaitGroup{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numGoroutines := r.Intn(1000) + 1

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go atomicCompareAndSwap(&lock, &wg)
	}

	wg.Wait()

	if atomic.LoadInt32(&lock) != 0 {
		t.Errorf("Expected lock to be 0, but got %d", lock)
	}
}
