package main

import "fmt"

func main() {

	quantity := 4
	length,width := 1.2,2.4
	customerName := "Karikalan"

	fmt.Println(customerName)
	fmt.Println("Ordered",quantity,"sheets of")
	fmt.Println(length*width,"sq units")

	fmt.Println("#Values of unassinged variables")
	
	var myInt int
	var myFloat int
	var myString string
	var myBool bool

	fmt.Println(myInt,myFloat,myString,myBool) 
}
