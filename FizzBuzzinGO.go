// package main

// import (
// 	"fmt"
// )

// func Main() {
// 	fmt.Print("Enter your number : ")
// 	var i int
// 	fmt.Scanln(&i)
// 	vat j:=1
// 	for j< 16; j++ {
// 		if j%i==0 {
// 			fmt.Printf("FizzBuzz",j )
// 		} else {
// 			fmt.Printf("%d", j)
// 		}
// 	}
// }

package main

import "fmt"

func main() {
	fmt.Print("Enter your number: ")

	var i int
	fmt.Scanln(&i)

	for j := 1; j < 16; j++ {
		if j%i == 0 {
			fmt.Printf("FizzBuzz ")
		} else {
			fmt.Printf("%d ", j)
		}
	}
	fmt.Println()
}
