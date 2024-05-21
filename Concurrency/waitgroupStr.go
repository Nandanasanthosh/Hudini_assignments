package main

import (
	"fmt"
	"time"
)

type Waitgroups struct {
	count int
}

func (wt *Waitgroups) Add() {
	wt.count += 1
}
func (wt *Waitgroups) Done() {
	wt.count -= 1
}
func (wt *Waitgroups) Wait() {
	for {
		if wt.count == 0 {
			break
		}
	}
}

func routine1s(wt *Waitgroups) {
	defer wt.Done()
	for i := 1; i < 6; i++ {
		fmt.Printf("%d in first program\n", i)
		time.Sleep(time.Millisecond)
	}
}

func routine2s(wt *Waitgroups) {
	defer wt.Done()
	for i := 1; i < 6; i++ {
		fmt.Printf("%d in second program\n", i)
		time.Sleep(time.Second)
	}
}

func routine3s(wt *Waitgroups) {
	defer wt.Done()
	for i := 1; i < 6; i++ {
		fmt.Printf("%d in third program\n", i)
		time.Sleep(time.Second)
	}
}
func main() {
	//waiterfn := new(sync.WaitGroup)
	wt := Waitgroups{
		count: 0,
	}
	wt.Add()
	go routine1s(&wt)
	wt.Add()
	go routine2s(&wt)
	wt.Add()
	go routine3s(&wt)

	wt.Wait()
}
