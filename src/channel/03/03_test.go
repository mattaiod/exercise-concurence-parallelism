package main

import (
	"sync"
	"testing"
)

func TestPipeline(t *testing.T) {
	numbers := make(chan int)
	squared := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go square(numbers, squared, &wg)

	go func() {
		for i := 0; i < 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	wg.Wait()

	expectedResults := []int{0, 1, 4, 9, 16}
	i := 0
	for result := range squared {
		if result != expectedResults[i] {
			t.Errorf("Expected %d, but got %d", expectedResults[i], result)
		}
		i++
	}
}
