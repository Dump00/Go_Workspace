package main

import (
	"fmt"
)

// strings are immutable

func main() {
	x := "Hello"
	y := "World"

	z := x + " " + y
	fmt.Println(z)
}
