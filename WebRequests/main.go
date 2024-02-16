package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to Golang Tut, we are studying handeling web requests")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %T\n", res)
	defer res.Body.Close()

	// reading
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	content := string(data)
	fmt.Println("The file is :", data)
	fmt.Println("The file is :", content)

}
