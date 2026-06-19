package main

import "fmt"

var orderNumber int
var drink string
var price int
var qty int

func showMenu(name string) {

	fmt.Println("")
	fmt.Println("============================")
	fmt.Println("1. Black Coffee - 120TK")
	fmt.Println("2. Latte - 150TK")
	fmt.Println("3. Cold Coffee - 180TK")
	fmt.Println("=============================")
	fmt.Printf("Type your Order number %s: ", name)

	fmt.Scan(&orderNumber)

	switch orderNumber {
	case 1:
		drink = "Black Coffee"
		price = 120
	case 2:
		drink = "Latte"
		price = 150
	case 3:
		drink = "Cold Coffee"
		price = 180
	default:
		fmt.Print("Please enter a valid number")
		price = 0
	}
	fmt.Printf("You Have Chosen %s", drink)
}

func quantity() int {
	fmt.Println("")
	fmt.Println("How many cups?: ")
	fmt.Scan(&qty)

	return qty
}

func receipt(customer string, drink string, price int, qty int) {

	var total = price * qty

	fmt.Println("===== RECEIPT =====")
	fmt.Printf("Customer: %s \n", customer)
	fmt.Printf("Drink: %s \n", drink)
	fmt.Printf("Quantity: %d \n", qty)
	fmt.Printf("Total: %d \n", total)
	fmt.Println("=======================")
}

func main() {
	var name string

	fmt.Print("Hey, What is your name: ")
	fmt.Scan(&name)

	fmt.Printf("Welcome, %s to Go cafe ☕", name)

	showMenu(name)
	quantity()
	receipt(name, drink, price, qty)
}
