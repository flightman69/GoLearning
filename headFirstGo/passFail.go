package main

import(
	"fmt"
	"bufio"
	"os"
	//"log"
	"strings"
	"strconv"
)

func main(){
	fmt.Print("Enter your Grade: ")
	reader := bufio.NewReader(os.Stdin)
	mark,_ := reader.ReadString('\n')
	mark = strings.TrimSpace(mark)
 	grade,_ := strconv.ParseFloat(mark,64)
	var status string
	if grade >= 60 {
		status = "Hooray!! You Have passed"
	} else {
		status = "Sorry You Have Failed"
	}
	fmt.Println("To pass the exam you atleast need 60%")
	fmt.Println("Your mark is over 60% ?: ",grade>=60)
	fmt.Println(status)
}
