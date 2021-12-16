package main

import "fmt"

func main() {
	// contador := 0
	// for i := 1; i <= 50; i++ {
	// 	if i%7 != 0 {
	// 		// fmt.Println(i)
	// 		continue
	// 	}
	// 	contador++
	// 	if contador == 4 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// Exercise 6

	age := -9

	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid age")
	case age < 18:
		fmt.Println("You are minor")
	case age == 18:
		fmt.Println("Congratulations!!")
	default:
		fmt.Println("You are major!")
	}
}
