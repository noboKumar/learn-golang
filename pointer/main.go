package main

import "fmt"

func changeValue(value *int) {
	*value = 1000
	fmt.Println("Inside Function: ", *value)
}

func main() {
	a := 20
	b := &a

	a = 100

	fmt.Println("a:", a)
	fmt.Println("b:", *b)
	fmt.Println("b:", b)
	
	
	inputValue := 20
	changeValue(&inputValue)

	fmt.Println("outside function", inputValue)
}


