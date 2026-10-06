// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var n int = 1000000
// 	var testSlice = []int{}
// 	var testslices2 = make([]int, 0, n)

// 	fmt.Println("Total time without preallocation:", timeLoop(testSlice, n))
// 	fmt.Println("Total time with preallocation: ", timeLoop(testslices2, n))
// }

// func timeLoop(slice []int, n int) time.Duration {
// 	var t0 = time.Now()
// 	for len(slice) < n {
// 		slice = append(slice, 1)
// 	}
// 	return time.Since(t0)
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	var n = 1000000
	var slice1 = []int{}
	var slcies2 = make([]int, 0, n)

	fmt.Println("append without preallocation:", timeloop(slice1, n))
	fmt.Println("append with preallocation:", timeloop(slcies2, n))
}

func timeloop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
