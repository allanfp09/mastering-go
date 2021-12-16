package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"

	score := []int{10, 20, 30}

	fmt.Printf("%d\n%f\n%s\n", x, y, z)
	fmt.Println("Gophers")
	fmt.Printf("%T\n%T\n", y, score)

	const b float64 = 1.422349587101

	fmt.Printf("%.4f", b)
}
