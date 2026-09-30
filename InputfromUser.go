package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // & = address of variable (pointers, later)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Printf("Hello %s, next year you'll be %d\n", name, age+1)
}
