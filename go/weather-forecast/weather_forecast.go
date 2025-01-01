// Package weather provides a simple weather forecast service.
package weather

// CurrentCondition is a global variable to store the current weather condition.
var CurrentCondition string

// CurrentLocation is a global variable to store the current location.
var CurrentLocation string

// Forecast function returns the weather forecast for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
