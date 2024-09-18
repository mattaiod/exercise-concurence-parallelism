package main

import "sync"

func selectMessages(ch1, ch2 <-chan string, ch3 chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case msg1 := <-ch1:
			ch3 <- msg1
		case msg2 := <-ch2:
			ch3 <- msg2
		default:
			close(ch3)
			return
		}
	}
}
