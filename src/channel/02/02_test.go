package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestProduce(t *testing.T) {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numProducers := r.Intn(1000) + 1
	expectedSum := 0

	for i := 0; i < numProducers; i++ {
		value := i
		expectedSum += value
		wg.Add(1)
		go produce(ch, value, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for val := range ch {
		sum += val
	}

	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}
}
