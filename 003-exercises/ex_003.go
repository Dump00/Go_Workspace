package main

import "fmt"

var e003_x int
var e003_y string
var e003_z bool

func main() {
	e003_x = 42
	e003_y = "James Bond"
	e003_z = true
	s := fmt.Sprintf("%v\t%v\t%v\t", e003_x, e003_y, e003_z)
	fmt.Println(s)
}
