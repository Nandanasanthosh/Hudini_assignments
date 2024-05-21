// Objective:
// Understand how to use channels for communication between goroutines.

// Task:

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:

// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.

// -------------------------------------------------------------------------------------
package main

import (
	"fmt"
	"sync"
	"time"
)

func numbers(c chan int, wt *sync.WaitGroup) {
	defer wt.Done()
	for i := 1; i < 11; i++ {
		c <- i
		time.Sleep(time.Second)
	}
}
func main() {
	channel := make(chan int)
	wt := new(sync.WaitGroup)
	wt.Add(1)
	go numbers(channel, wt)
	for i := 1; i < 11; i++ {
		fmt.Print(<-channel)
	}
	wt.Wait()
	close(channel)
}
