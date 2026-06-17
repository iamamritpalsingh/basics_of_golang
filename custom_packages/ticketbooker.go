package my_custom_package

import "strconv"

var BookingObject = make([]map[string]string, 0) // List of maps

func BookTicketForUser(first_name string, last_name string, user_tickets int, email string, city string) {
	// This func is just to understand the map data structure
	var map_object = make(map[string]string)
	map_object["name"] = first_name + " " + last_name
	map_object["email"] = email
	map_object["city"] = city
	map_object["number_of_tickets"] = strconv.Itoa(user_tickets)
	BookingObject = append(BookingObject, map_object)
}