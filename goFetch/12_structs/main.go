package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)

}

// hasBirthday method (ponter receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (ponter receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "George", lastName: "Reyes", city: "Austin", gender: "M", age: 22}

	fmt.Println(person1)
	fmt.Println(person1.firstName)

	// Increments the age by 1 unit
	person1.hasBirthday()

	// If female, will change their last name to value
	person1.getMarried("Reyes")
	fmt.Println(person1.greet())
}
