package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"time"
	"math/rand"
	"strconv"
)

func main(){
	seed := time.Now().Unix()
	rand.Seed(seed)
	target := rand.Intn(100)
	
	success := false
	for x:= 0; x<5; x++ {
		fmt.Printf("You have %d chances left\nEnter your Guess: ",5-x)
		reader := bufio.NewReader(os.Stdin)
		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess,_ := strconv.Atoi(input)

		if guess > target {
			fmt.Println("You've Guessed High")
		} else if guess < target {
			fmt.Println("You've Guess Low")
		}	else {
			success = true
			fmt.Println("Hooray You've guessed right")
			break
		}
}

if success == false{
	fmt.Printf("You've lost the number was %d\n",target) 
	}
}
