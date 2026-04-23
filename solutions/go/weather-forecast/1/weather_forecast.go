// Package weather provides outputs for the current weather conditions for a city.
package weather

var (
	// CurrentCondition reprents current weather conditions as a string.
	CurrentCondition string
	// CurrentLocation represents a city as a string.
	CurrentLocation string
)

// Forecast takes two strings, city and condition and outputs another string indicating the conditions at that location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
