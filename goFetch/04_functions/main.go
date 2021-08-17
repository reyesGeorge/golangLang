package main

import "fmt"

// You want to add the type of the input and also the type of the return
func greeting(name string) string {
	return "Hello" + name

}

func getSum(num int, num2 int) int {
	return num + num2
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(greeting("George!"))
}
