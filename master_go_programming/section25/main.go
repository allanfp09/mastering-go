package main

import (
	"fmt"
)

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func main() {
	// Exercise 1
	var m money = 6.7 * 4.5566
	fmt.Println(m)
	m.print()

	// Exercise 2
	s := m.printStr()
	fmt.Println(s)
}
