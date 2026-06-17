package main

import "strings"

func getFirstNames(bookings[] string) []string {
	var first_names_slice = []string{}

	for _, value := range bookings {
		var names = strings.Fields(value)
		first_names_slice = append(first_names_slice, names[0])
	}
	return first_names_slice
}