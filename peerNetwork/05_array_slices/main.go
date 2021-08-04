package main

import "fmt"

func main() {
	// You have to specify the length and type of the Array in go
	// Here we are specifying a string array that can hold 2 values
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declaring and Assigning at the same time
	fruitArray := [2]string{"Apple", "Orange"}

	fruitArrayy := []string{"Apple", "Orange", "Bannana"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])

	fmt.Println(fruitArray)
	fmt.Println(fruitArrayy)

}
