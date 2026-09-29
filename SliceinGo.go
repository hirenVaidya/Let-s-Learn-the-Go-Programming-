//Dynamic list that can grow and shrink in size. Slices are more common than arrays in Go. Slices are built on top of arrays and provide a more powerful interface to sequences of data.

package main

import "fmt"

func main() {
	array := []string{"Hiren", "Ashok"}
	array = append(array, "Vaidya")
	fmt.Println(array)

	prices := []float64{19.99, 50.5, 3.75}
	total := 0.0

	for _, price := range prices {
		total += price
		//	fmt.Println(price)
	}
	fmt.Println("Total:", total)
}
