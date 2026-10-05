package main

import (
	"errors"
	"fmt"
)

func main() {
	printfunction()
	var n int = 20
	var m int = 0
	var result, rem, err int = addingfunction(n, m, err)
	if err != 0 {
		fmt.Println(err.Error())
		return
	} else if rem == 0 {
		fmt.Println("The division is exact, result is %v", result)
	}
	//fmt.Println(result)
	//fmt.Println(rem)
	fmt.Println("Hello division %v and remainder %v", result, rem)
}
func printfunction() {
	fmt.Println("Hello World")
}

func addingfunction(n int, m int) (int, int, err) {
	var err error = 0
	if m == 0 {
		err = errors.New("Division by zero is not allowed")
		return 0, 0, err
	}
	return n / m, n % m, err
}
