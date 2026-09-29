package main

import "fmt"

func main() {
	a := "There is no place like 127.0.0.1"
	primes := [6]int{2, 3, 4, 5, 6, 7}
	var s []int = primes[1:4]
	fmt.Println("Hello world")
	fmt.Println(a)
	fmt.Println(primes)
	fmt.Println(s)
	ages := make(map[string]int)

	ages["John"] = 25
	ages["Alice"] = 30
	ages["Bob"] = 22

	fmt.Println(ages["John"])
}
