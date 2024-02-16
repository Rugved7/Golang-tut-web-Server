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
	DecodeJSON()
}

func DecodeJSON() {
	jsonData := []byte(`
	{
		"Name": "C++",
		"Price": 499,
		"Platform": "Rugved.com",
		"Passwords": "adc71877",
		"Tags": [
			"DSA ",
			"Game Developement"
		]
	}
    `)

	// Decoding This JSON data
	var RugvedCourse course
	isValid := json.Valid(jsonData) // check validity of json data

	if isValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonData, &RugvedCourse)
		fmt.Printf("%#v\n", RugvedCourse)
	} else {
		fmt.Println("JSON data is not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}
}
