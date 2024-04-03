package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainingTickets int = 50
var bookings = make([]userData, 0)

type userData struct{
	firstName string
	lastName string
	email string
	numberOfTickets int
}

func main() {


	greetUsers()
	
	for remainingTickets > 0 && len(bookings) < 50{

		firstName, lastName, email, userTickets := getUserImput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserImput(firstName, lastName, email, userTickets, remainingTickets)


		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
	

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v.\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short ")
			}
			if !isValidEmail {
				fmt.Println("Email adress you entered does not contain @ sign ")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entereid is invalid")
			}
			
		}	
	}
	
}

func greetUsers () {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


func getUserImput() (string, string, string, int) {
	var firstName string
		var lastName string
		var email string
		var userTickets int
		//Ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter your number of tickets: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}

func bookTicket (userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for user

	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v.\n", bookings)


	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a conformation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	println("######")
	fmt.Printf("Sending ticket:\n %v \nto email adress %v\n", ticket, email)
	println("######")

}
