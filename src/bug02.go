package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // Global wait group

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	ch := make(chan int)

	wg.Add(1) // Add 1 to the counter

	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait() // Wait until counter is 0. Waits until Print has finished executing.

}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	wg.Done() // Decreases counter by 1. Lets the wait group know that the wait is over.
}
