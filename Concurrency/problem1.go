package main

import "fmt"

// Objective:
// Learn how to send and receive values using channels.

// Task:
func Helloworld(s chan string) {

	s <- "hello world"
}

// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
// The main function should receive the message from the channel and print it.
func main() {
	msg := make(chan string)
	// Helloworld := make(chan string)
	// go func(){ Helloworld <-"Hello, World!"}()
	go Helloworld(msg)
	message := <-msg
	fmt.Println(message)
}

// Hints:

// Use an unbuffered channel for simplicity.

// -------------------------------------------------------------------------------------

/



// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.
