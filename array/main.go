package main

import "fmt"

func main() {
	var arr1 = []int{1, 2, 3, 4, 5}
	arr1 = append(arr1, 200)
	fmt.Println(arr1)
	fmt.Println(cap(arr1))
	fmt.Println(len(arr1))
}
