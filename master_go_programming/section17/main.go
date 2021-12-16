package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// Exercise 1
	// file, err := os.Create("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// Exercise 2
	// _, err := os.Stat("data.txt")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Fatal("Sorry, this file does not exist.")
	// 	}
	// }

	// err = os.Rename("info.txt", "data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Exercise 3
	// err = os.Remove("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Exercise 4
	// r, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%v", r)

	// Exercise 5
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// Exercise 6
	// message := []byte("The Go gopher is an iconic mascot!")
	// err := ioutil.WriteFile("info.txt", message, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// My own
	// message := []byte("Hi i am Allan and i like golang")
	// file, err := os.OpenFile("info.txt", os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	file, err := os.Open("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Sorry this file has not been created")
		}
	}
	defer file.Close()

	// err = os.WriteFile("info.txt", message, 06044)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Bytes()
	fmt.Println(text)

}
