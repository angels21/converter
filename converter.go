package main

import "fmt"

type Converter struct{ Centimeter, Feet, Minutes, Seconds, Kilograms, Pounds, Fahrenheit, Celsius float64 }

func (cvr Converter) CentimeterToFeet() float64 {
	return cvr.Feet * 30.48
}

func (cvr Converter) FeetToCentimeter() float64 {
	return cvr.Centimeter * 0.0328084
}
func (cvr Converter) MinutesToSeconds() float64 {
	return cvr.Minutes * 60
}
func (cvr Converter) SecondsToMilliseconds() float64 {
	return cvr.Seconds * 1000
}
func (cvr Converter) PoundsToKilograms() float64 {
	return cvr.Pounds * 0.454
}
func (cvr Converter) KilogramsToPounds() float64 {
	return cvr.Kilograms * 2.205
}
func (cvr Converter) FahrenheitToCelsius() float64 {
	return ((cvr.Fahrenheit - 32) * (9 / 5))
}
func (cvr Converter) CelsiusToFahrenheit() float64 {
	return ((cvr.Celsius * (9 / 5)) + 32)
}
func main() {

	cvr := Converter{3, 5, 8, 6, 5, 10, 67, 35}
	fmt.Println(cvr.Feet, "feet is: ", cvr.CentimeterToFeet(), "centimeter")
	fmt.Println(cvr.Centimeter, "centimeter is: ", cvr.FeetToCentimeter(), "feet")
	fmt.Println(cvr.Minutes, "Minutes is: ", cvr.MinutesToSeconds(), "seconds")
	fmt.Println(cvr.Seconds, "Seconds is: ", cvr.SecondsToMilliseconds(), "Milliseconds")
	fmt.Println(cvr.Pounds, "Pounds is: ", cvr.PoundsToKilograms(), "Kilograms")
	fmt.Println(cvr.Kilograms, "Kilograms is: ", cvr.KilogramsToPounds(), "Pounds")
	fmt.Println(cvr.Fahrenheit, "Fahrenheit is: ", cvr.FahrenheitToCelsius(), "Celsius")
	fmt.Println(cvr.Celsius, "Celsius is: ", cvr.CelsiusToFahrenheit(), "Fahrenheit")

}
