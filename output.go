package main

import "fmt"

func main() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)   // Prints the value in the default format
	fmt.Printf("%#v\n", i)  // Prints the value in Go-syntax format
	fmt.Printf("%v%%\n", i) // Prints the % sign
	fmt.Printf("%T\n", i)   // Prints the type of the value

	fmt.Printf("%v\n", txt)  // Prints the value in the default format
	fmt.Printf("%#v\n", txt) // Prints the value in Go-syntax format
	fmt.Printf("%T\n", txt)  // Prints the type of the value
}
