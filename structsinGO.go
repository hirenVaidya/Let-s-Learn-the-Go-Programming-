package main

import "fmt"

type ganEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesleft() uint8 {
	return e.mpkwh * e.kwh
}

func (e ganEngine) milesleft() uint8 {
	return e.mpg * e.gallons
}

type engine interface {
	milesleft() uint8
}

func canmakeit(e engine, miles uint8) {
	if miles <= e.milesleft() {
		fmt.Println("Yes, you can make it there")
	} else {
		fmt.Println("Need to fuel up first ")
	}
}

func main() {
	myEngine := ganEngine{25, 15}
	canmakeit(myEngine, 50)
}
