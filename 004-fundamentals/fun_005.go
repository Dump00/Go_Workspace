package main

import "fmt"

// Iota - pre declared identifier

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
