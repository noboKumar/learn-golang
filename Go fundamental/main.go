package main

import "fmt"

func makeCoffee(drinkType string) string {
	coffee := fmt.Sprintf("making %s Coffee", drinkType)
	return coffee
}

func makeTea (teaType string)(tea string, price int){
	price = 20
	tea = fmt.Sprintf("I am having a %s Tea and it's %d Tk", teaType, price)
	return tea, price
}

func coffeeOrder(customerName string, coffeeType string, price int)(orderData string) {
	orderData = fmt.Sprintf("Order for %s: %s coffee costs %d Taka", customerName, coffeeType, price)
	return orderData
}

func main() {
	// var name = "nobo"
	// var name string = "nobo"

	// name := "nobo"

	// var (
	// 	name = "nobo"
	// 	age  = 20
	// )

	// var name string = "Nobo"
	// name  = "dev"

	// fmt.Println(name)

	// var name string = "Nobo"
	// var age int = 20

	// fmt.Printf("Hi, My name is %s and i am %d years old.", name, age)

	// makeCoffee("Black")
	fmt.Println(coffeeOrder("nobo", "Black", 50))
}

