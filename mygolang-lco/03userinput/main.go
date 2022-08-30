package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Movie")

	// comma ok || comma err
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your rating ", input)
	fmt.Printf("The type of rating is %T\n", input)
}
