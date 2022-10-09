package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)
	fmt.Println(reflect.TypeOf(myInt),reflect.TypeOf(myIntPointer))
	myBool := false
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer) 

	myFloat := 4
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	
}
