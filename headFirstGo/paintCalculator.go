package main

import (
	"fmt"
	"log"
)

func main(){
	var area, tot float64
	area,err := paintNeeded(2.3,5.6)
	tot += area
	area,err = paintNeeded(4.5,5.6)
	tot += area
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Total: %.2f litres of paint needed\n",tot)
}

func paintNeeded(height float64, width float64)(float64,error){
	if height < 0 || width < 0{
		return 0, fmt.Errorf("Either height or width is invalid")
	} 
	area := height *  width
	return area/10.0,nil
}


