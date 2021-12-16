package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func cube(v float64) float64 {
	return math.Pow(v, 3)
}

func f1(n uint) (int, int) {
	fac := 1
	var sum int

	for i := 1; i <= int(n); i++ {
		fac *= i
		sum++
	}
	return fac, sum
}

func myFunc(m string) int {
	a, err := strconv.Atoi(m)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := strconv.Atoi(m + m)

	c, _ := strconv.Atoi(m + m + m)

	return a + b + c
}

func sum(a ...int) int {
	s := 0

	for _, v := range a {
		s += v
	}
	return s
}

func sum1(a ...int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func searchItem(st []string, s string) bool {
	for _, v := range st {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	// Exercise 1
	r := cube(3)
	fmt.Printf("%.f\n", r)

	// Exercise 2
	f, s := f1(4)
	fmt.Println(f, s)

	// Exercise 3
	fmt.Println(myFunc("3"))

	// Exercise 4
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// Exercise 5
	fmt.Println(sum(1, 2, 3, 4, 5, 6))

	// Exercise 6
	slice := searchItem([]string{"leon", "tiger"}, "tiger")
	fmt.Println(slice)
}
