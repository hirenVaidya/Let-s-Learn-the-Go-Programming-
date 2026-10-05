//Dynamic list that can grow and shrink in size. Slices are more common than arrays in Go. Slices are built on top of arrays and provide a more powerful interface to sequences of data.

package main

import "fmt"

func main() {
	// array := []string{"Hiren", "Ashok"}
	// array = append(array, "Vaidya")
	// fmt.Println(array)

	// prices := []float64{19.99, 50.5, 3.75}
	// total := 0.0

	// for _, price := range prices {
	// 	total += price
	// 	//	fmt.Println(price)
	// }
	// fmt.Println("Total:", total)

	intArr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(intArr)

	var slice []int = []int{4, 5, 6}
	fmt.Println(slice)
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))
	slice = append(slice, 7)
	fmt.Println(slice)
	fmt.Println("Length of slice after append:", len(slice))
	fmt.Println("Capacity of slice after append:", cap(slice))

}
