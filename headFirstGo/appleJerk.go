package main

import "fmt"

func main() {
	var totApples = 10
	var jerkAte = 4 
	var applesLeft = totApples - jerkAte

	fmt.Println("I started with",totApples,"apples.")
	fmt.Println("Some jerk ate",jerkAte,"apples.")
	fmt.Println("There are",applesLeft,"left.")
}
