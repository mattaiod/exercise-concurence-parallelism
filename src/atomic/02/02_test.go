package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomicSetBool(t *testing.T) {
	var flag int32
	wg := sync.WaitGroup{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numGoroutines := r.Intn(1000) + 1

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go atomicSetBool(&flag, 1, &wg)
	}

	wg.Wait()

	if atomic.LoadInt32(&flag) != 1 {
		t.Errorf("Expected flag to be 1, but got %d", flag)
	}
}
