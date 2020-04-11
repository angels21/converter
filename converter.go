package main

import (
	"fmt"
	"math"
)

//Converter struct for various units
type Converter struct{ Centimeter, Feet, Minutes, Seconds, Kilograms, Pounds, Degree, Radian, Fahrenheit, Celsius float64 }

//CentimeterToFeet converter
func (cvr Converter) CentimeterToFeet() float64 {
	return cvr.Feet * 30.48
}

//FeetToCentimeter converter
func (cvr Converter) FeetToCentimeter() float64 {
	return cvr.Centimeter * 0.0328084
}

//MinutesToSeconds converter
func (cvr Converter) MinutesToSeconds() float64 {
	return cvr.Minutes * 60
}

//SecondsToMilliseconds converter
func (cvr Converter) SecondsToMilliseconds() float64 {
	return cvr.Seconds * 1000
}

//PoundsToKilograms converter
func (cvr Converter) PoundsToKilograms() float64 {
	return cvr.Pounds * 0.454
}

//KilogramsToPounds converter
func (cvr Converter) KilogramsToPounds() float64 {
	return cvr.Kilograms * 2.205
}

//DegreeToRadian converter
func (cvr Converter) DegreeToRadian() float64 {
	return (cvr.Degree * 2 * math.Pi) / (360)
}

//RadianToDegree converter
func (cvr Converter) RadianToDegree() float64 {
	return (cvr.Radian * 360) / (2 * math.Pi)
}

//FahrenheitToCelsius converter
func (cvr Converter) FahrenheitToCelsius() float64 {
	return ((cvr.Fahrenheit - 32) * (9 / 5))
}

//CelsiusToFahrenheit converter
func (cvr Converter) CelsiusToFahrenheit() float64 {
	return ((cvr.Celsius * (9 / 5)) + 32)
}
func main() {

	cvr := Converter{152.4, 5, 8, 6, 4.54, 10, 360, 6.286, 67, 35}
	fmt.Println(cvr.Feet, "feet is: ", cvr.CentimeterToFeet(), "centimeter")
	fmt.Println(cvr.Centimeter, "centimeter is: ", cvr.FeetToCentimeter(), "feet")
	fmt.Println(cvr.Minutes, "Minutes is: ", cvr.MinutesToSeconds(), "seconds")
	fmt.Println(cvr.Seconds, "Seconds is: ", cvr.SecondsToMilliseconds(), "Milliseconds")
	fmt.Println(cvr.Pounds, "Pounds is: ", cvr.PoundsToKilograms(), "Kilograms")
	fmt.Println(cvr.Kilograms, "Kilograms is: ", cvr.KilogramsToPounds(), "Pounds")
	fmt.Println(cvr.Degree, "Degrees is: ", cvr.DegreeToRadian(), "Radians")
	fmt.Println(cvr.Radian, "Radians is: ", cvr.RadianToDegree(), "Degrees")
	fmt.Println(cvr.Fahrenheit, "Fahrenheit is: ", cvr.FahrenheitToCelsius(), "Celsius")
	fmt.Println(cvr.Celsius, "Celsius is: ", cvr.CelsiusToFahrenheit(), "Fahrenheit")

}
