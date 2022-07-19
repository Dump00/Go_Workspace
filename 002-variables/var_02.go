/* Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
Println formats using the default formats for its operands and writes to standard output.
Spaces are always added between operands and a newline is appended.
It returns the number of bytes written and any write error encountered.
*/

package main

import "fmt"

func main() {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old.")
}
