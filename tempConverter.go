package main

import "fmt"

type celsius float64
type fahrenheit float64
type calctemp func()

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func displayFahrenheit() {
	for i := -40; i <= 100; i = i + 5 {
		var value celsius = celsius(i)
		fmt.Printf("| %-10.1f | %-10.1f |\n", value, value.fahrenheit())
	}
}

func displayCelsius() {
	for i := -40; i <= 212; i = i + 9 {
		var value fahrenheit = fahrenheit(i)
		fmt.Printf("| %-10.1f | %-10.1f |\n", value, value.celsius())
	}
}

func displayTable(c calctemp, col1, col2 string) {
	fmt.Printf("===========================\n")
	fmt.Printf("| %-10v | %-10v |\n", col1, col2)
	fmt.Printf("===========================\n")
	c()
	fmt.Printf("===========================\n")
}

func main() {
	displayTable(displayFahrenheit, "°C", "°F")
	displayTable(displayCelsius, "°F", "°C")
}
