package main

import (
	"fmt"
	"reflector"
)

func main(){
	no := 6
	fmt.Printf("%T\n",no)
	fmt.Printl(reflector.TypeOf(n))
}
