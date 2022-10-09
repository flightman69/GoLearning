package main

import (
	"fmt"
)

func main(){
	var myInt int = 10
	var myIntPointer *int = &myInt
	fmt.Printf("myIntPointer = %p\n",myIntPointer)
	fmt.Printf("&myInt = %p\n",&myInt)
	fmt.Printf("myInt = %d\n",myInt)
	*myIntPointer = 12
	fmt.Printf("After *myIntPointer = 12\n")
	fmt.Printf("myInt = %d\n",myInt)
} 
