package main

import "fmt"

func main() {
	var y int = 30
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	var z = "Hello"
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	greet := `hello "go" developers`
	fmt.Println(greet)
}
