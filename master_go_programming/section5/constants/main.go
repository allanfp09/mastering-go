package main

import "fmt"

func main() {
	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)

	// Exercise 3

	const (
		secPerDay = (60 * 60) * 24
		daysYear  = 365
	)
	fmt.Printf("%d", secPerDay*daysYear)
}
