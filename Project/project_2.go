// Objective:Create a command-line interface (CLI) application in Go that allows users to input a block of text and then calculates the frequency of each word in the text. This project will help you understand and implement basic Go concepts like variables, loops, conditionals, maps, functions, and strings.
// Requirements:
// Input Text: Allow users to input a block of text.
// Process Text: Count the frequency of each word in the text.
// Display Frequencies: Display the word frequencies in a readable format.
// Exit: Allow the user to exit the application.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Count(text string) {
	t := strings.TrimSpace(text)
	word := strings.Split(strings.ToLower(t), " ")
	fmt.Println(word)
	count := make(map[string]int)
	for _, val := range word {
		if _, exists := count[val]; !exists {
			count[val] = 1
		} else {
			count[val] += 1
		}
	}
	fmt.Println(count)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input text :")
	input, _ := reader.ReadString('\n')
	Count(input)
	fmt.Print("Do you want to exit(y/n)")
	in, _ := reader.ReadString('\n')
	if in == "y" {
		os.Exit(0)
	}

}
