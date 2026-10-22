package main

import "fmt"

func main() {
	fmt.Print("Enter your number : ")
	var n int
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Printf("The number %d is even.", n)
	} else {
		fmt.Printf("The number %d is odd.", n)
	}
}
