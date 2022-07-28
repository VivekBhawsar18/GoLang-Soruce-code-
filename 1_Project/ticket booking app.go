package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings:= []string{}

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for{

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name : ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name : ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your e-mail : ")
		fmt.Scan(&email)

		fmt.Println("Enter no of tickets : ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings =append(bookings ,firstName+" "+ lastName)

		fmt.Printf("The whole array: %v\n" , bookings)
		fmt.Printf("The first value: %v\n" , bookings[0])
		fmt.Printf("Array type : %T\n" , bookings)
		fmt.Printf("Array length: %v\n" , len(bookings))
	

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation e-mail at %v\n" , firstName , lastName , userTickets , email)
		fmt.Printf("%v tickets are remaining for %v\n" , remainingTickets , conferenceName)
	}

}
