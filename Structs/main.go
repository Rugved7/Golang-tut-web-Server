package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	Rugved := User{"Rugved", "rugved@go.dev", 20, "Male", false}
	fmt.Println(Rugved)
	fmt.Printf("The Type of the var Rugved is : %+v\n", Rugved)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Gender string
	Status bool
}
