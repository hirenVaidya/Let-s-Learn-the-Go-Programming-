package main

import "fmt"

func main() {
	var i int = 42
	var i8 int8 = 127
	var u uint = 100
	var b byte = 255

	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	var isGoFun bool = true

	var s string = "Go is fun"

	var r rune = 'A'

	var cx complex128 = complex(1, 2)

	fmt.Println(i, i8, u, b)
	fmt.Println(f32, f64)
	fmt.Println(isGoFun, s)
	fmt.Println(r, string(r))
	fmt.Println(cx)
}
