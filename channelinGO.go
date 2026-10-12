package main

import "fmt"

func main() {
	var c = make(chan int)
	go Process(c)
	for i := range c {
		fmt.Println(i)
	}
}
func Process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	//close(c)
}
