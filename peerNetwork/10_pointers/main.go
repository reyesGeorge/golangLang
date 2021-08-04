package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)

	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b, "star B")
	fmt.Println(*&a, "star amp A")

	// Use the * to get the value of the address
	// Use the & to get the address of the variable

	// Change val with pointer
	*b = 10
	fmt.Println(a)

}
