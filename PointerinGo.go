package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Println("The Value P point to is :", *p)
	fmt.Println("The Value of i is :", i)
	p = &i
	*p = 2
	fmt.Println("The Value P point to is :", *p)
	fmt.Println("The Value of i is :", i)

	var k int32 = 3
	i = k
	fmt.Println("The Value of k is :", k)
	fmt.Println("The Value of i is :", i)

}
