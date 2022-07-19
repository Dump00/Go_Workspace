package main

import "fmt"

func main() {
	type hotdog int
	var b hotdog
	b = 12
	fmt.Println(b)
	fmt.Printf("%T", b)
}
