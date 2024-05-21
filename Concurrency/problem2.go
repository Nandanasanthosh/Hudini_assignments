package main

import (
	"fmt"
	"sync"
	"time"
)

// Objective:
// Learn how to create and use goroutines.

// Task:

// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:

// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.

// -------------------------------------------------------------------------------------
func routine1(wt *sync.WaitGroup) {
	defer wt.Done()
	for i := 1; i < 6; i++ {
		fmt.Printf("%d in first program\n", i)
		time.Sleep(time.Millisecond)
	}
}
func routine2(wt *sync.WaitGroup) {
	defer wt.Done()
	for i := 1; i < 6; i++ {
		fmt.Printf("%d in second program\n", i)
		time.Sleep(time.Second)
	}
}
func routine3(wt *sync.WaitGroup) {
	defer wt.Done()
	for i := 1; i < 6; i++ {
		fmt.Printf("%d in third program\n", i)
		time.Sleep(time.Second)
	}
}
func main() {
	waiterfn := new(sync.WaitGroup)
	waiterfn.Add(3)

	go routine1(waiterfn)
	go routine2(waiterfn)
	go routine3(waiterfn)

	waiterfn.Wait()
}
