package main

import "fmt"

func main() {
	var myInt int = 2
	var myFloat float64 = 3.4
	fmt.Println("Type of myInt is %T and myFloat is%T",myInt,myFloat)
	fmt.Println("#After Conversion: ")

	var myInt = float64(myInt)
	fmt.Println("Type of myInt is %T",myInt)
	
}
