package main

import "fmt"

func main() {
	var inslices = []int{1, 2, 3, 4, 5}
	fmt.Println(suminslices(inslices))
	var floatinslices = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(suminslices(floatinslices))
	var float64inslices = []float64{1.11, 2.22, 3.33, 4.44, 5.55}
	fmt.Println(suminslices(float64inslices))

}

// func suminslices[T int | float32 | float64](s []T) T {
// 	var sum T
// 	for _, v := range s {
// 		sum += v
// 	}
// 	return sum
// }
func suminslices[T any](s []T) bool {
	return len(s) == 0
}
