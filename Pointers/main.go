package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var ptr *int
	fmt.Println("value of pointer is", ptr)

	number := 18
	var ptr2 = &number
	fmt.Println("The value of the pointer to number is", ptr2)
	fmt.Println("The value of the pointer to number is", *ptr2)
}
