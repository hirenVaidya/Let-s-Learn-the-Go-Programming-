package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter your numberthat you need to sum : ")
	i := 0
	fmt.Scanln(&i)
	fmt.Printf("The sum of numbers from 1 to %d is: %d", i, sumfornumber(i))
}
func sumfornumber(n int) int {
	sum := 0
	i := 1
	for i <= n {
		sum += i
		i++
	}
	return sum
}
