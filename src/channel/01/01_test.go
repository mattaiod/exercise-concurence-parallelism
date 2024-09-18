package main

import (
	"sync"
	"testing"
)

func TestSendMessage(t *testing.T) {
	ch := make(chan string)
	wg := sync.WaitGroup{}

	message := "Hello, Go!"
	wg.Add(1)
	go sendMessage(ch, message, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	received := <-ch
	if received != message {
		t.Errorf("Expected '%s', but got '%s'", message, received)
	}
}
