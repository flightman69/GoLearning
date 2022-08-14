package main

import (
	"fmt"
	"strings"
)

// Show Details
var showName string = "The Lan Party"

const tot_tickets uint = 50

var remainingTickets = tot_tickets
var bookings = []string{}

func main() {

	for {
		//Greet Users
		greetUsers()

		//Get User Input
		userName, userTickets := getUserInput()
		if userTickets > remainingTickets || userTickets == 0 {
			fmt.Printf("Invalid amount of tickets only %v tickets left\n", remainingTickets)
			continue
		}
		fmt.Printf("Thank you %v for buying %v tickets\nHave a Great day :)\n\n", userName, userTickets)

		//Update Booking and remaining details
		bookings, remainingTickets = bookingUpdate(bookings, remainingTickets, userName, userTickets)

		//Checks if all tickets sold out
		if remainingTickets == 0 {
			//end of program
			fmt.Print("All tickets sold out :(\n\n")
			fmt.Print("===============================================================\n\n")
			break
		}
	}

}

func greetUsers() {
	fmt.Print("===============================================================\n")
	fmt.Printf("\t\tWelcome to the %v \n\t\tYou can buy your tickets here\n\t\tOnly %v tickets remaining!\n", showName, remainingTickets)
	fmt.Print("===============================================================\n\n")

}

func getUserInput() (string, uint) {

	fmt.Println("What's your name ?")
	var userName string
	fmt.Scanf("%v", &userName)

	fmt.Printf("Hello %v how many tickets you like to buy ??\n", userName)
	var userTickets uint
	fmt.Scan(&userTickets)

	return userName, userTickets

}

//geto booking details

func getBookingDetails(totBookings []string) []string {
	firstNames := []string{}
	for _, names := range totBookings {
		var first_name = strings.Split(names, "_")
		firstNames = append(firstNames, first_name[0])
	}

	return firstNames

}

func bookingUpdate(bookings []string, remaining_tickets uint, userName string, userTickets uint) ([]string, uint) {
	//Updating booking and tickets details
	bookings = append(bookings, userName)
	remaining_tickets = remaining_tickets - userTickets

	//Print booking details
	bookingDetails := getBookingDetails(bookings)
	fmt.Printf("Total Bookings %v and now %v tickets are remaining\n", bookingDetails, remaining_tickets)
	fmt.Print("===============================================================\n\n")

	return bookings, remaining_tickets
}
