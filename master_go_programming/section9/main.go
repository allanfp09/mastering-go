package main

import "fmt"

func main() {
	var cities [2]string
	fmt.Printf("%#v\n", cities)

	grades := [3]float64{45.6, 58.4}
	fmt.Printf("%#v\n", grades)

	taskDone := [...]bool{
		true, false, true, false, false,
	}
	fmt.Printf("%#v\n", taskDone)

	nums := [...]int{30, -1, -6, 90, -6}

	count := 0
	for _, v := range nums {
		if v > 0 {
			count++
		}
	}
	fmt.Println(count)
}
