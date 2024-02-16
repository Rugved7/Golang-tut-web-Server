package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string
	Price     int
	Platform  string
	Passwords string
	Tags      []string
}

func main() {
	fmt.Println("Hello, world!")
	EncodeJSON()
}

func EncodeJSON() {
	myCourses := []course{
		{"Golang", 499, "Rugved.com", "adc777", []string{"Web dev ", "Backend Developement"}},
		{"C++", 499, "Rugved.com", "adc71877", []string{"DSA ", "Game Developement"}},
	}
	// Make this Raw data as JSON
	makeJSON, err := json.Marshal(myCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Json format: %s\n", makeJSON)

	JSON, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Json format: %s\n", JSON)
}
