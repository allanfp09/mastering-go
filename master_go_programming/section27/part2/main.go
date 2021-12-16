package main

import "fmt"

func power(v int, ch chan int) {
	ch <- v * v
}

func main() {
	// Exercise 1

	// var c1 float64

	// c2 := make(chan<- rune)
	// c3 := make(<-chan rune)
	// c4 := make(chan int, 10)

	// fmt.Printf("%T\n", c1)
	// fmt.Printf("%T\n", c2)
	// fmt.Printf("%T\n", c3)
	// fmt.Printf("%T\n", c4)

	// Exercise 2
	// c := make(chan string)

	// go func(s string) {
	// 	c <- s
	// }("I love Go!")

	// r := <-c
	// fmt.Println(r)

	// Exercise 3
	// c := make(chan int)

	// go func(n int) {
	// 	c <- n
	// }(100)

	// fmt.Println(<-c)

	// Exercise 4

	ch := make(chan int)

	// for i := 1; i <= 50; i++ {
	// 	go power(i, ch)
	// }

	// for i := 1; i <= 50; i++ {
	// 	fmt.Println(<-ch)
	// }

	// Exercise 4
	for i := 1; i <= 50; i++ {
		go func(v int) {
			ch <- v * v
		}(i)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-ch)
	}
}
