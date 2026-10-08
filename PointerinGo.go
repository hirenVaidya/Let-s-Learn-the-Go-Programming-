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

	fmt.Println("-------The End-------")
	fmt.Println("--Pointer in slices--")

	var slice = []int{1, 3, 4}
	var slice2 = slice
	slice2[2] = 5
	fmt.Println("The Value of slice is :", slice)
	fmt.Println("The Value of slice is :", slice2)

}
