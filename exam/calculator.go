package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ArgumentError = fmt.Errorf("Needed more arguments for operation")

// Calculator function to perform operations based on the operand
func Calculator(value1, value2 float64, operand string) (float64, error) {
	// TODO: Implement Calculator function
	var res float64
	switch operand {
	case "+":
		res = (value1 + value2)
	case "-":
		res = (value1 - value2)
	case "*":
		res = (value1 * value2)
	case "/":
		if value2 == 0 {
			return -1, errors.New("Can't divide be zero\n")
		} else {
			res = (value1 / value2)
		}
	default:
		return -1, errors.New("Invalid operation.")
	}
	return res, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter the expression (e.g., 10 + 20): ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}

		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		// TODO: Check if exact 3 parts are given. If not, throw an error
		if len(parts) < 3 {
			fmt.Errorf("Error :%w\n", ArgumentError)
		}
		// TODO: Take all 3 values and pass to calculator function
		num1, _ := strconv.ParseFloat(parts[0], 64)
		num2, _ := strconv.ParseFloat(parts[2], 64)
		result, errorOut := Calculator(num1, num2, parts[1])
		if errorOut != nil {
			fmt.Printf("Error : %v\n", errorOut)
		}
		// TODO: Print results
		fmt.Printf("Result: %v\n", result)
	}

}
