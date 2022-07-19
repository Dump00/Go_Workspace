// var keyword

package main

import "fmt"

var z = "meow!"

// zero value
var a int

func main() {
	var y = 25
	fmt.Println(y)
	fmt.Println(a)
	foo2()
	fmt.Println(a)
}

func foo2() {
	fmt.Println(z)
	a = 12
}
