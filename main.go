package main

import "fmt"

func main() {

	// fmt.Print("hello fellaas")
	// fmt.Println("hello fellaas")

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	// fmt.Println("Welcome to our conference booking application")
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

}
