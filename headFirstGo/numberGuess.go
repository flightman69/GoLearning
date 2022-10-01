/*
#		Generate a random number from 1 - 100
#		Prompt the player to guess the number
#		if the player guess lowed display "Low"
#		if the player guess high display "High"
#		Allow the player guess upto 10 times
# 	If the player guessed right displahy "Won"
#		If player ran out of chances display "Lost"
#	 	Finally display the number "
*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100)
	fmt.Print("The Go God has choosen a number between 1 -100\n\n")
	success := false
	for x := 0; x < 10; x++ {
		fmt.Print("You have ", 10-x, " chances to meet the GOd\n\n")
		fmt.Print("Enter your guess: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, _ := strconv.Atoi(input)
		fmt.Println(guess)

		if guess == target {
			success = true
			fmt.Println("Hooray you have guessed the number correctly :)")
			break
		} else if guess > target {
			fmt.Print("You guessed high\n\n")
		} else {
			fmt.Print("You guessed low\n\n")
		}
	}
	if !success {
		fmt.Print("O.Ops!! you ran out of your chances\nNow burn in hell")
		fmt.Println("The GOd choosen number was: %v", target)
	}
}
