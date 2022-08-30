package main

import "fmt"

const ConstVariable string = "cnst"

func main() {
	var username string = "Flightman"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T\n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type : %T\n", smallval)

	var smallfloat float32 = 25.23432523
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type : %T\n", smallfloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type : %T\n", anotherVariable)

	//implicit type
	var website = "gnu.org"
	fmt.Printf("This is an implicit declaration and it's value and type is %v and %T\n", website, website)

	// no var style
	numberOfUser := 30000
	fmt.Printf("This variable is assined using ':=' Its value is %v\n", numberOfUser)

	fmt.Printf("This is an global variable : %v\n", ConstVariable)

}
