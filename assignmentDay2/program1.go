package main

import (
	"math"
	"strings"
)

func main() {
	monkeyCount(17)
	MakeNegative(18)
	FindMultiples(2, 8)
	CountBy(2, 9)
	PowersOfTwo(7)
	Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"})
	GetSum(-2, 2)
	FindShort("bitcoin take over the world maybe who knows perhaps")
	SquareSum(([]int{1, 2, 8}))
	OddCount(10938490)
	SortMyString("Dinner")
	LeastLarger([]int{4, 2, 6, 2, 6}, 2)
}

func monkeyCount(n int) []int {
	// Your code here, happy coding!
	var arr = []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i+1)
	}
	return arr
}

func MakeNegative(x int) int {
	switch {
	case x > 0:
		return -x
	case x < 0:
		return x
	default:
		return 0
	}
}

func FindMultiples(integer, limit int) []int {
	s := integer
	var arr = []int{}
	for i := 2; s <= limit; i++ {
		arr = append(arr, s)
		s = integer * i
	}
	return arr
}

func CountBy(x, n int) []int {
	var arr = []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i*x)
	}
	return arr
}

func PowersOfTwo(n int) []uint64 {
	var arr = []uint64{}
	for i := 0; i <= n; i++ {
		final_value := math.Pow(2, float64(i))
		arr = append(arr, uint64(final_value))
	}
	return arr
}

func Points(games []string) int {
	xPoints := 0
	for i := 0; i < 10; i++ {
		x := strings.Split(games[i], ":")
		if x[0] > x[1] {
			xPoints += 3
		} else if x[0] == x[1] {
			xPoints += 1
		}
	}
	return xPoints
}

func GetSum(a, b int) int {
	first, second := a, b
	if b > a {
		first = b
		second = a
	}
	sum := 0
	for i := second; i <= first; i++ {
		sum += i
	}
	return sum
}

func FindShort(s string) int {
	word := strings.Split(s, " ")
	maxLen := 1000

	for i := 0; i < len(word); i++ {
		if len(word[i]) < maxLen {
			maxLen = len(word[i])
		}
	}
	return maxLen
}

func SquareSum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i] * numbers[i]
	}
	return sum
}

func OddCount(n int) int {
	return n / 2 //eg : 5/2 =2(1,3) 2 odd cases
}

func SortMyString(s string) string {
	even, odd := "", ""
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			even = even + string(s[i])
		} else {
			odd = odd + string(s[i])
		}
	}
	return even + " " + odd
}

func LeastLarger(a []int, i int) int {
	index := 0
	smallest := 1000
	for y := 0; y < len(a); y++ {
		if a[y] > a[i] {
			if a[y] < smallest {
				smallest = a[y]
				index = y
			}
		}
	}
	if smallest == 1000 {
		return -1
	} else {
		return index
	}
}
