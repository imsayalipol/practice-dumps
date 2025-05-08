package main

import (
	"fmt"
)

func main() {

	var f func(int) int
	f = func(x int) int {
		return x * x
	}
	fmt.Println(f(5)) // Output: 25

	// you can also assign name
	f = square
	fmt.Println(f(10))
}

func square(i int) int {
	return i * i
}
