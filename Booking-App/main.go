package main

import "fmt"

func main() {
	show_name := "Go Show"
	const tot_show_tickets uint = 50
	var remaining_tickets = 50
	fmt.Printf("tot_show_tickets is %T, show_name is a %T\n", tot_show_tickets, show_name)
	fmt.Printf("Welcome to %v booking application\n", show_name)
	fmt.Println("Get your tickets here to attend! Only", remaining_tickets, "tickets left")

	var user_name string
	var user_tickets int
	// ask user for their name

	user_name = "Mike"
	user_tickets = 3
	fmt.Printf("User %v booked %v tickets\n", user_name, user_tickets)
}
