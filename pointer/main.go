package main

import "fmt"

func main() {
	a := 20
	b := &a

	a = 100

	fmt.Println("a:", a)
	fmt.Println("b:", *b)
	fmt.Println("b:", b)
}
