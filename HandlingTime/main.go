package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")

	Time := time.Now()
	fmt.Println(Time)

	fmt.Println(Time.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2003, time.July, 07, 06, 36, 00, 00, time.UTC)
	fmt.Println(createDate)
}
