//Dynamic list that can grow and shrink in size. Slices are more common than arrays in Go. Slices are built on top of arrays and provide a more powerful interface to sequences of data.

package main

import "fmt"

func main() {
	array := []string{"Hiren", "Ashok"}
	array = append(array, "Vaidya")
	fmt.Println(array)
}
