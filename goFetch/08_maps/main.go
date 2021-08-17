package main

import "fmt"

func main() {

	// Define a Map
	emails := make(map[string]string)

	// Second way to declare and add values in one go
	emails2 := map[string]string{"Bob": "bob@gmail.com", "George": "george@gmail.com"}

	emails["Bob"] = "bob@gmail.com"
	emails["George"] = "george@gmail.com"

	fmt.Println(emails)

	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Deleteing a Value from a Maps
	// delete(emails, "Bob")

	fmt.Println(len(emails2))
}
