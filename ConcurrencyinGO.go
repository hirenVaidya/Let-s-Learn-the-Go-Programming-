package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var data = []string{"Hiren", "Ashok", "Vaidya", "Alka", "Neha"}

func main() {
	t0 := time.Now()
	for i := 0; i < len(data); i++ {
		go datatansfer(i)
	}
	fmt.Println("total time taken:", time.Since(t0))
}
func datatansfer(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The Result from the database is:", data[i])
}
