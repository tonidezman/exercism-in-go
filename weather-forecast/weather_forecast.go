// Package weather is a package.
package weather

// CurrentCondition is a var.
var CurrentCondition string

// CurrentLocation is also a var.
var CurrentLocation string

// Forecast -> is a function.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
