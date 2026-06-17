package my_custom_package

import (
	"time"
)

// Custom type to store different type of data types
type TicketDetails struct {
	FirstName string
	TimeOfBooking time.Time
}
var AllTicketDetails = make([]TicketDetails, 0)

func CollectTicketDetails(first_name string) []TicketDetails {
	var TicketData = TicketDetails	{
		FirstName : first_name,
		TimeOfBooking: time.Now(),
	}
	AllTicketDetails = append(AllTicketDetails, TicketData)
	return AllTicketDetails
}
