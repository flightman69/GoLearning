package main

import (
	"fmt"
)

func createPointer() *float64 {
	myFloat := 86.8
	return &myFloat
}
func main(){
	myFloatPointer := createPointer()
	fmt.Printf("Value: %f lives in the address %p\n",*myFloatPointer,myFloatPointer)
}	
