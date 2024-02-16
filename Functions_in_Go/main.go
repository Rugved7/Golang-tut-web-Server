package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions module")
	greet()
	addFunction(7, 18)
}

func greet() {
	fmt.Println("Namaste Rugved")
}

func addFunction(a int, b int) {
	result := a + b
	fmt.Println(result)
}
