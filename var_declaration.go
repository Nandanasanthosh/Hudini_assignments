package main

import "fmt"

var c, python, java bool

func main() {
	var i int = 9
	var ch, num, str, boolean = 'v', 90, "is string", true
	//ch wont give character value but it will return unicode value
	fmt.Println(i, boolean, ch, num, str, boolean)
}
