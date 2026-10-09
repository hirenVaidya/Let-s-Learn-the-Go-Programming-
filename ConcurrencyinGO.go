package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var data = []string{"Hiren", "Ashok", "Ranjan", "Neha", "Falguni"}
var result = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go datatansfer(i)
	}
	wg.Wait()
	fmt.Println("total time taken:", time.Since(t0))
	fmt.Println("The Result are", result)
}
func datatansfer(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The Result from the database is:", data[i])
	result = append(result, data[i])
	wg.Done()
}
