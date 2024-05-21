package main

import (
	"errors"
	"fmt"
)

// Stack represents a stack that holds a slice of integers.
type Stack struct {
	elements []int
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
	// Appends the new element to stack element
}

// Pop removes the element at the top of the stack and returns it.
// It returns an error if the stack is empty.
func (s *Stack) Pop() (int, error) {
	n := len(s.elements)
	if n == 0 {
		return 0, errors.New("stack is empty")
	} else {
		poppedElement := s.elements[n-1]
		s.elements = append(s.elements[:n-1], s.elements[n:]...)
		return poppedElement, nil
	}

}

// Peek returns the element at the top of the stack without removing it.
// It returns an error if the stack is empty.
func (s *Stack) Peek() (int, error) {
	n := len(s.elements)
	if n == 0 {
		return 0, errors.New("stack is empty")
	} else {
		return s.elements[n-1], nil
	}

}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *Stack) IsEmpty() bool {
	n := len(s.elements)
	if n == 0 {
		return true
	} else {
		return false
	}
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	n := len(s.elements)
	return n
}

func main() {
	stack := Stack{}

	// Example usage (You can modify this part to test your implementation)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack size:", stack.Size())

	topElement, err := stack.Peek()
	if err == nil {
		fmt.Println("Top element:", topElement)
	} else {
		fmt.Println(err)
	}

	for !stack.IsEmpty() {
		element, err := stack.Pop()
		if err == nil {
			fmt.Println("Popped element:", element)
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println("Stack size after popping all elements:", stack.Size())
}
