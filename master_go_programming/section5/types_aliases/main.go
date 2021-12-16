package main

import "fmt"

type duration int
type mile float64
type kilometer float64

const m2km = 1.609

func main() {
	var hour duration
	fmt.Printf("%d\t%T\n", hour, hour)
	hour = 3600
	fmt.Printf("%d", hour)

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Printf("%f", kmBerlinToParis)

}
