package main

import "fmt"

type person struct {
	name  string
	age   int
	color []string
	gd    grades
}

type grades struct {
	grade  float64
	course string
}

func main() {
	// Exercise 1
	me := person{
		name:  "Allan",
		age:   22,
		color: []string{"red", "blue"},
		gd:    grades{89.65, "Math"},
	}
	you := person{
		name:  "Joel",
		age:   45,
		color: []string{"yellow", "brown"},
		gd:    grades{90.00, "Programming"},
	}

	//Exercise 2
	me.name = "Andrei"

	var colors []string = you.color

	colors = append(colors, "green")

	you.color = colors

	fmt.Printf("%T, %v\n", colors, colors)
	fmt.Printf("This you struc: %v\n", you)

	// Exercise 3
	for i, c := range me.color {
		fmt.Println(i, c)
	}

	// Exercise 4
	me.gd.grade = 98.0
	me.gd.course = "Golang"

	fmt.Println(me)
}
