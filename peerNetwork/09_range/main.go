package main

import "fmt"

func main() {
	ids := []int{33, 32, 333, 445, 64, 13, 1, 23, 34}

	// Using an Index
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using and Index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	// Range with Map
	emails := map[string]string{"Bob": "bob@gmail.com", "George": "george@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name " + k)
	}

}
