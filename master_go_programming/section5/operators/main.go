package main

import (
	"fmt"
)

func main() {
	var (
		i int     = 3
		f float64 = 3.2
		// c              = 3
		// g              = 3.2
		// s1, s2         = "3.14", "5"
	)

	convertI := float64(i)
	convertF := int(f)

	fmt.Printf("%.1f\n%d", convertI, convertF)

	// Exercise 2

	// convC := fmt.Sprintf("C is:", c)
	// convG := fmt.Sprintf("G is:", g)
	// convS2, err := strconv.Atoi(s2)
	// _ = err
	// convS1, err := strconv.ParseFloat(s1, 64)
	// _ = err

	// fmt.Printf("%s\n%T\n", convC, convC)
	// fmt.Printf("%s\n%T\n", convG, convG)
	// fmt.Printf("%d\n%T\n", convS2, convS2)
	// fmt.Printf("%.2f\n%T\n", convS1, convS1)

	// Exercise 4

	var time int64

	time = (149_600_000 * 1000) / 299_792_458
	fmt.Printf("%d", time)
}
