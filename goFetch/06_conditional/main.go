package main

import (
	"fmt"
)

func main() {

	x := 15
	y := 10

	// If else
	if x < y {
		fmt.Printf("%d is more than %d\n", y, x)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("the color is blue")
	} else {
		fmt.Println("The color is neither blue nor red")
	}

	// Switch statement
	switch color {
	case "red":
		fmt.Println("The color is red")
	case "blue":
		fmt.Println("The color is blue")
	default:
		fmt.Println("The color is neither blue nor red")
	}
}
