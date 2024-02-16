package main

import "fmt"

func main() {
	name := "Rugved"
	fmt.Println(name)

	var friends [3]string

	friends[0] = "Mohit"
	friends[1] = "Sai"
	friends[2] = "Rugved"

	fmt.Println(friends)
	fmt.Println(len(friends))

	var Friends = [3]string{"Rugved", "Mohit", "Sai"}
	fmt.Println("The Friends Array looks like this", Friends)

	// Slices

	var college = []string{}
	fmt.Printf("Type of college is %T\n", college)

	var skills = []string{"React", "Nodejs", "C++"}
	fmt.Println("Skills:=", skills)

	skills = append(skills, "Golang", "Next.js")
	fmt.Println("Skills:=", skills)

	// skills = append(skills[1:3])
	fmt.Println("Skills:", skills)

}
