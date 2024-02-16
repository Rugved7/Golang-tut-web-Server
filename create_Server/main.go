package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	PerformGETRequest()
	PerformPOSTRequest()
	PerformFormData()
}

func PerformGETRequest() {
	const myURL = "http://localhost:8000/get"

	res, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status: ", res.Status)
	defer res.Body.Close()

	// Printing String using string Builder

	var resString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	bytecount, _ := resString.Write(content) // for giving count of bytes in the string
	fmt.Println("Content length : ", bytecount)
	fmt.Println("Content is:", resString.String())

	// fmt.Println("Content: ", string(content))
}

func PerformPOSTRequest() {
	const URL = "http://localhost:8000/post"
	// Generate Fake JSON Data
	resBody := strings.NewReader(`
{
	"name": "Rugved",
    "age": 25,
    "friends": ["Mohit", "Sai", "Rugved"]
}
`)
	response, err := http.Post(URL, "application/json", resBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

// Form Data
func PerformFormData() {
	const MYurl = "http://localhost:8000/postform"
	// Fake Form Data
	data := url.Values{}
	data.Add("Name", "Rugved")
	data.Add("Surname", "Agasti")
	data.Add("Age", "20")

	res, err := http.PostForm(MYurl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
