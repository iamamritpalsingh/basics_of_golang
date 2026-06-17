package my_custom_package

import (
	"fmt"
	"time"	
	"math/rand"
)

var Session_number = rand.Intn(100000) // Global variable
func ShowTime() {
	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")
	fmt.Println("Logged in time:", formattedTime)
}