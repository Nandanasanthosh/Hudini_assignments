package main

import "fmt"

//1st program
func CalculateAverage(numbers []int) float64 {
	sum := 0
	// Write your code here to calculate and return the average of the array elements.
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return float64(sum / len(numbers))
}

//2nd program
func checkAge(age int) string {
	// Write your code here to determine and print if the age corresponds to a minor, a young adult, or an adult.
	switch {
	case age <= 18:
		return "minor"
	case age >= 27:
		return "adult"
	case age > 18, age < 27:
		return "young adult"
	default:
		return "Enter correct value"
	}
}

//3rd program
func reverseString(str string) string {
	var reverse string = ""
	// Write your code here to reverse and return the string.
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return reverse
}

//4
func findLargestNumber(numbers1 []int) int {
	largest := 0
	// Write your code here to find and return the largest number in the array.
	for i := 0; i < len(numbers1); i++ {
		if numbers1[i] > largest {
			largest = numbers1[i]
		}
	}
	return largest
}

//5
/*
type Counter struct{
	countn int
	}
func (count Counter) add() Counter {
	return count.countn++
}
func (count Counter) subtract() Counter {
	return count.countn--
}
func (count Counter) display() Counter {
	fmt.println(count.countn)
}
func (count Counter) reset() Counter {
	count.countn=0
	return count.countn
}*/
	// Example usage:
  // Output: 0
// Example usage:
func main() {
	//1
	str := []int{10, 20, 30, 40, 50}
	fmt.Printf("CalculateAverage({10, 20, 30, 40, 50})= %d \n", CalculateAverage(str))
	//2
	fmt.Printf("Checkage(25)= %s\n", checkAge(25))
	//3
	strl := "hello"
	fmt.Printf("reverseString(hello)= %s\n", reverseString(strl))
	//4
	fmt.Printf("largest({10, 20, 30, 40, 50})= %d", findLargestNumber(str))
	var counter =Counter{countn:0}
	
	counter=counter.add();
	counter.add();
	counter=counter.display();  // Output: 2
	counter=counter.subtract();
	counter.display();  // Output: 1
	counter.reset();
	counter.display();
}

type Counter struct{
	countn int
	}
func Counterfn(count Counter) Counter{
	func add(count Counter) counter{
		return count.countn++
	}
	func subtract(count Counter) counter{
		return count.countn--
	}
	func (count Counter) display() Counter {
		fmt.println(count.countn)
	}
	func (count Counter) reset() Counter {
		count.countn=0
		return count.countn
	}		
}