// Objective:
// Understand how to handle multiple senders with a single receiver using channels.

// Task:

// Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// The main function should receive both messages from the channel and print them.

// -------------------------------------------------------------------------------------

package main

import (
	"fmt"
	"time"
	//"sync"
)

func msg1(c chan string) {
	c <- "Hello"
}
func msg2(c chan string) {
	time.Sleep(time.Second)
	c <- "World"
}

func main() {
	ch := make(chan string)
	//wt := new(sync.WaitGroup)
	go msg1(ch)
	go msg2(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
