package main

import (
	"fmt"
	"strings"
)

func main() {
	// var name = "Allan"

	// country := "Panama"

	// fmt.Printf(`My name is %s\n`, name)
	// fmt.Printf("My country is %s\n", country)

	// fmt.Println("Hi: \"Hello\"")
	// fmt.Println("C:\\Users\\a.txt")

	// Exercise 2
	// r := 'ă'
	// fmt.Printf("%T\n", r)

	// s1, s2 := "ma", "m"
	// fmt.Println(s1 + s2 + string(r))

	// Exercise 3
	s1 := "țară means country in Romanian"

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v\t\n", s1[i])
	}

	fmt.Println(strings.Repeat("#", 10))

	for i, v := range s1 {
		fmt.Printf("index: %d, rune: %c\n", i, v)
	}

}
