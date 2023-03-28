package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName              string
	lastName               string
	email                  string
	numberOfTickets        uint
	isOptedInForNewsletter bool
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// for {
	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Conference tickets has been booked out. We hope to see you come back next year.")
			// break
		}
	} else if userTickets == remainingTickets {
		// do something about this
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}

		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}

		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}

	}

	wg.Wait()

	// }

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	// _ - blank identifier
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// var firstName = names[0]
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

// func validateUserInput(firstNameRef string, lastNameRef string, emailRef string, userTicketsRef uint) (bool, bool, bool) {
// 	isValidName := len(firstNameRef) >= 2 && len(lastNameRef) >= 2
// 	isValidEmail := strings.Contains(emailRef, "@")
// 	isValidTicketNumber := userTicketsRef > 0 && userTicketsRef <= remainingTickets

// 	return isValidName, isValidEmail, isValidTicketNumber
// }

func getUserInput() (string, string, string, uint) {
	var firstNameRef string
	var lastNameRef string
	var emailRef string
	var userTicketsRef uint

	fmt.Println("enter first name")
	fmt.Scan(&firstNameRef)

	fmt.Println("enter last name")
	fmt.Scan(&lastNameRef)

	fmt.Println("enter email address")
	fmt.Scan(&emailRef)

	fmt.Println("enter tickets count")
	fmt.Scan(&userTicketsRef)

	return firstNameRef, lastNameRef, emailRef, userTicketsRef
}

func bookTicket(fName string, lName string, emailRef string, userTicketsRef uint) {
	remainingTickets = remainingTickets - userTicketsRef

	// var myslice []string
	// var mymap map[string]string

	// var userData = make(map[string]string)
	// userData["firstName"] = fName
	// userData["lastName"] = lName
	// userData["email"] = emailRef
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTicketsRef), 10)

	var userData = UserData{
		firstName:              fName,
		lastName:               lName,
		email:                  emailRef,
		numberOfTickets:        userTicketsRef,
		isOptedInForNewsletter: true,
	}

	// bookings = append(bookings, fName+" "+lName)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will receive a confirmation email at %v.\n", fName, lName, userTicketsRef, emailRef)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTicketsR uint, firstNameR string, lastNameR string, emailR string) {
	// fmt.Printf("%v tickets for %v %v", userTickets, firstName, lastName)
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicketsR, firstNameR, lastNameR)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, emailR)
	fmt.Println("####################")
	wg.Done()
}
