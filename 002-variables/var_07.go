package main

import "fmt"

func main() {
	var a int
	type hotdog int
	var b hotdog
	b = 12
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)
}
