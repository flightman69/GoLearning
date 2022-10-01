package main

import (
	"fmt"
	"os"
	"log"
	//"bufio"
	)

func main(){
	//fmt.Print("Enter the file name: ")
	//reader := bufio.NewReader(os.Stdin)
	//fileName,_ := reader.ReadString('\n')
	fileInfo,err := os.Stat("tax.go")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}	
