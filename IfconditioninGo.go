package main

import "fmt"

func main() {
	var stock int
	fmt.Print("Enter stock quantity: ")
	fmt.Scanln(&stock)
	if stock == 0 {
		fmt.Println("Out of stock")
	} else if stock < 20 {
		fmt.Println("Running low on stock")
	} else {
		fmt.Println("Stock is available")
	}
}
