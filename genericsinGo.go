package main

import "fmt"

func main() {
	var inslices = []int{1, 2, 3, 4, 5}
	var floatinslices = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	var float64inslices = []float64{1.11, 2.22, 3.33, 4.44, 5.55}
	fmt.Println(suminslices(inslices))
	fmt.Println(Intsuminslices(inslices))
	fmt.Println(float32suminslices(inslices))
	fmt.Println(float64suminslices(inslices))
}

func suminslices(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}
func Intsuminslices(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}
func float32suminslices(s []float32) float32 {
	var sum float32
	for _, v := range s {
		sum += v
	}
	return sum
}
func float64suminslices(s []float64) float64 {
	var sum float64
	for _, v := range s {
		sum += v
	}
	return sum
}
