package main

import (
	"fmt"
	"booking-app/custom_packages" // booking-app is the module name and custom_packages is the package name
	"time"
	"sync"
)
// Package level variable, these are accessbile throughout the file
var bookapp_name = "Bookmyshow" 
var bookapp_address = "Mumbai, Maharashtra, India" 
var vg = sync.WaitGroup{}

func main() {
	const total_tickets = 100
	remaining_tickets := 100 // different way to define variable

	my_custom_package.ShowTime()
	fmt.Println("Current session number is:", my_custom_package.Session_number)
	greetUsers(total_tickets, remaining_tickets)
	// Define all user inputs
	var first_name string
	var last_name string
	var email string
	var user_tickets int
	var bookings = []string{}
	var city string
	// for {
		//Get user input
		fmt.Println("Enter your first name")
		fmt.Scan(&first_name)
		fmt.Println("Enter your last name")
		fmt.Scan(&last_name)
		fmt.Println("Enter your email address")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets")
		fmt.Scan(&user_tickets)
		fmt.Println("Enter city")
		fmt.Scan(&city)

		if remaining_tickets < user_tickets {
			fmt.Printf("You cannot book %v ticktes, we have only %v tickets left.\n", user_tickets,remaining_tickets)
			// break;
		}
		bookings = append(bookings, first_name+" "+last_name)
		remaining_tickets = remaining_tickets - user_tickets

		var first_names_slice = getFirstNames(bookings)

		if (remaining_tickets == 0) {
			fmt.Println("Show sold out")
			// break;
		}
		fmt.Printf("User %v has booked %v tickets. Confirmation will be sent to %v\n", first_name+" "+last_name, user_tickets, email)
		fmt.Printf("After booking %v tickets, now %v tickets are left out\n", user_tickets, remaining_tickets)
		fmt.Printf("Whole slice %v\n", first_names_slice)
		my_custom_package.BookTicketForUser(first_name, last_name, user_tickets, email, city)
		fmt.Printf("Booked tickets map: %v\n", my_custom_package.BookingObject)

		var my_struct = my_custom_package.CollectTicketDetails(first_name)
		fmt.Printf("Structure data is %v\n", my_struct)
		go sendEmail(email, user_tickets)
		vg.Add(1) // Here 1 as we have only one go routine
		vg.Wait()
		// This has nothing to do with the booking, it is just a example of switch case
		// switch city {
		// 	case "Australia","New Zealand":
		// 		fmt.Println("You booked your tickets in Pacific ocean area.")
		// 	case "Germany":
		// 		fmt.Println("You booked in europe.")	
		// 	default:
		// 		fmt.Println("You booked invalid ticket.")	
		// 		return
		// }
	// }
}

func greetUsers(total_tickets int, remaining_tickets int) {
	fmt.Printf("Welcome to %v %v booking application\n", bookapp_name, bookapp_address)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", total_tickets, remaining_tickets)
}

func sendEmail(email string, tickets int) {
	fmt.Println("Sending email thread started")
	time.Sleep(10 * time.Second) // 10 seconds wait/sleep
	var content = fmt.Sprintf("Email sent to %v for %v tickets", email, tickets)
	fmt.Printf("%v", content)
	vg.Done()
}
