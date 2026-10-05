package main

import "fmt"

func main() {
	printfunction()
	var n int = 20
	var m int = 10
	var result int = addingfunction(n, m)
	fmt.Println(result)
}

func printfunction() {

	fmt.Println("Hello World")
}

func addingfunction(n int, m int) int {
	return n / m

}
