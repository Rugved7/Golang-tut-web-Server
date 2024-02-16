package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file section in Golang")
	content := "This is the default content for this section"

	// Making a file
	file, err := os.Create("./file.txt")

	// if we encounter an error
	if err != nil {
		panic(err) // if  err found , return error
	}

	// else we gonna write the content in the file
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote", length, "characters to the file")
	defer file.Close()

	// Read file
	readFile("/run/media/not_rugved/HDD/GoLang/file.txt")
}

// Reading a file
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside file is:\n", databyte)
}
