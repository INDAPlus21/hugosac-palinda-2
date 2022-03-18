package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)

	go func() { // Add the string to the channel in another goroutine
		ch <- "Hello world!"
	}()

	fmt.Println(<-ch) // The value can now be accessed
}
