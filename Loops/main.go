package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Loops module in Go")

	days := []string{"Sun", "Mon", "Tue", "Wed"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	for j := range days {
		fmt.Println(days[j])
	}
}
