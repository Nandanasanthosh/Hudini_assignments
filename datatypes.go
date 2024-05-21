/*
	    PROBLEM: NOT RUNNING
		# command-line-arguments
		.\datatypes.go:5:2: could not import math/cmplx (open : The system cannot find the file specified.)
*/
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Tobe   bool       = false
	Maxint uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt()(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T Value: %v\n", Maxint, Maxint)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//Constants cannot be declared using the := syntax.
	var i int
	const world ="nin"
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %v %q\n", i, f, b,world, s)

	//unsigned type conversion:
	var x, y int = 3, 4
	i :=-23
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(i)
	fmt.Println(i,x, y,f, z)
}

}
