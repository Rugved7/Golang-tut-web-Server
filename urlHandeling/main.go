package main

import (
	"fmt"
	"net/url"
)

const Myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj"

func main() {
	fmt.Println("Hello from Golang")
	fmt.Println(Myurl)

	// Parsing
	result, err := url.Parse(Myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)

	fmt.Println(result.Query())
}
