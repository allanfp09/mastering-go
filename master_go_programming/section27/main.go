package main

import (
	"fmt"
	"math"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func sum(a, b float64, wg *sync.WaitGroup) {
	fmt.Printf("%.2f\n", a+b)
	wg.Done()
}

func main() {
	// Exercise 1
	var wg sync.WaitGroup
	wg.Add(50)
	// go sayHello("Mr. Wick", &wg)
	// wg.Wait()

	// Exercise 2
	// go sum(4.5, 5.6, &wg)
	// go sum(3.4, 5.0, &wg)
	// go sum(6, 4, &wg)

	// Exercise 3
	// go func(v float64, wg *sync.WaitGroup) {
	// 	fmt.Printf("The square value of %v is %v", v, math.Sqrt(v))
	// 	wg.Done()
	// }(9, &wg)

	// Exercise 4
	for i := 100; i < 150; i++ {
		go func(v float64, wg *sync.WaitGroup) {
			fmt.Printf("%v\n", math.Sqrt(v))
			wg.Done()
		}(float64(i), &wg)
	}

	wg.Wait()
}
