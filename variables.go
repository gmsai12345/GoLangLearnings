package main

import "fmt"

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

// The Print() function prints its arguments with their default format.
func main() {
	// go variables
	var numbers int = 10
	nums := 15
	fmt.Println(numbers)
	fmt.Println(nums)
	var a string
	var b int
	var c bool
	// ASSIGNED TO default values
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	var student1 string
	student1 = "John"
	fmt.Println(student1)
	//Can be used inside and outside of functions	Can only be used inside functions
	//Variable declaration and value assignment can be done separately	Variable declaration and value assignment cannot be done separately (must be done in the same line)
	var aa, bb, cc, dd int = 1, 3, 5, 7

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	var (
		a1 int
		b1 int    = 1
		c1 string = "hello"
	)

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	const PI = 3.14
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	var i, j string = "Hello", "World"
	// adding space between strings
	fmt.Print(i, " ", j)
	var x, y = 10, 20

	fmt.Print(x, y)
}
