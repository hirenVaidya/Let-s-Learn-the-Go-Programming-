package main

import "fmt"

func main() {
	fmt.Print("Enter your first number: ")
	var a int
	fmt.Scanln(&a)
	fmt.Print("Enter your second number")
	var b int
	fmt.Scanln(&b)
	fmt.Print("Enter your third number: ")
	var c int
	fmt.Scanln(&c)
	switch {
	case a > b && a > c:
		fmt.Printf("The largest number is %d", a)
	case b > a && b > c:
		fmt.Printf("The largest number is %d", b)
	case c > a && c > b:
		fmt.Printf("The largest number is %d", c)
	case a == b && a == c:
		fmt.Printf("All numbers are equal")
	}
}
