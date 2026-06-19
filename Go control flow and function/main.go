package main

import "fmt"

// func makeCoffee() {
// 	fmt.Printf("Hot Coffee...")
// }

func main() {
	// makeCoffee := func() {
	// 	fmt.Printf("Hot Coffee...")
	// }
	// makeCoffee()

	func() {
		fmt.Printf("Hot Coffee...")
	}()
}
