package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int)
	go Process(c)
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}
func Process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Existing Process")
	//close(c)
}
