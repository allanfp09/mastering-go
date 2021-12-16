package main

import "fmt"

func main() {
	// Exercise 1
	// x := 10.10
	// fmt.Println(&x)

	// pr := &x
	// fmt.Printf("%T, %p, %v", pr, pr, *pr)

	// Exercise 2
	// x, y := 10, 2
	// prtx, prty := &x, &y
	// z := *prtx / *prty
	// fmt.Printf("%T, %v", z, z)

	// Exercise 3
	x, y := 5.5, 8.8
	x1, y1 := &x, &y
	*x1, *y1 = 8.8, 5.5
	fmt.Println(x, y)
}
