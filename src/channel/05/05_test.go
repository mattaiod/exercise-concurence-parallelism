package main

import (
	"sync"
	"testing"
)

func TestSelectMessages(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go selectMessages(ch1, ch2, ch3, &wg)

	go func() {
		ch1 <- "from ch1"
		close(ch1)
	}()

	go func() {
		ch2 <- "from ch2"
		close(ch2)
	}()

	wg.Wait()

	expectedMessages := []string{"from ch1", "from ch2"}
	i := 0
	for msg := range ch3 {
		if msg != expectedMessages[i] {
			t.Errorf("Expected '%s', but got '%s'", expectedMessages[i], msg)
		}
		i++
	}
}
