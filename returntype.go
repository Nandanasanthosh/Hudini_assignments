package main

import "fmt"

//it is a naked return. It returns all the naked values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//to take float32 values
func splitfloat(sum float32) float32 {
	var avg = sum / 4
	return avg
}
func main() {
	fmt.Println(split(17))
	fmt.Println(splitfloat(12.3))
}
