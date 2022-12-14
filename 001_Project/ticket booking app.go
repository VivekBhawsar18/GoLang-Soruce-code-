package main

import (
	"fmt"
	"strings"
)

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

		if userTickets <= remainingTickets {
		
			remainingTickets = remainingTickets - userTickets
			bookings =append(bookings ,firstName+" "+ lastName)

			firstNames := []string{}
			for _, booking := range bookings{

				var names = strings.Fields(booking)
				firstNames := append(firstNames , names[0])
			}
			fmt.Printf("The first names of bookings are %v\n" , firstNames)
			
			fmt.Printf("The whole array: %v\n" , bookings)
			fmt.Printf("The first value: %v\n" , bookings[0])
			fmt.Printf("Array type : %T\n" , bookings)
			fmt.Printf("Array length: %v\n" , len(bookings))
			
		
			fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation e-mail at %v\n" , firstName , lastName , userTickets , email)
			fmt.Printf("%v tickets are remaining for %v\n" , remainingTickets , conferenceName)
			
			if remainingTickets == 0 {
				
				fmt.Println("Our conffernce is booked out . Come back next year . ")
				break
			}
		} else{
			fmt.Printf("We have only %v ticktes remaining  , so you can't book %v tickets\n" , remainingTickets , userTickets)
		}


	}

}
