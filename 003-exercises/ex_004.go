package main

import "fmt"

type cisco int

var e004_x cisco

func main() {
	fmt.Println(e004_x)
	fmt.Printf("%T\n", e004_x)

	e004_x = 42

	fmt.Println(e004_x)
}
