package main

import (
	// Standard Libs
	"fmt"
	"strings"

	// Third Party Libs
	"cron/service"
)

// comments added
func main() {
	var choice string
	fmt.Println("Do you want to get the weather data using city or pincode?")
	fmt.Scan(&choice)

	switch strings.ToLower(choice) {
	case "city":
		var city string
		fmt.Println("Enter a city name:")
		fmt.Scan(&city)
		service.GetWeatherByCity(city)

	case "pincode":
		var pincode string
		fmt.Println("Enter a pincode:")
		fmt.Scan(&pincode)
		service.GetWeatherByPincode(pincode)

	default:
		return
	}
}
