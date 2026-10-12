package main

import "fmt"

func main() {
	var c = make(chan int)
}

func Process(c chan int) {
	c <- 123

}

