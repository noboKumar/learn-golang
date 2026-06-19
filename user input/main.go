package main

import "fmt"

func main() {
	var userInput string

	fmt.Print("What is your name: ")
	fmt.Scan(&userInput)

	fmt.Printf("Hello, %s How are you?", userInput)
}
