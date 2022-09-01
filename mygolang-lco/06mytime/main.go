package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createDate := time.Date(2020, time.April, 21, 00, 00, 00, 00, time.UTC)

	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
