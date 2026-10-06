package main

import "fmt"

type ganEngine struct {
	mpg       uint8
	gallons   uint8
	ownerinfo owner
	int
}

type owner struct {
	name string
}

func main() {
	myEngine := ganEngine{25, 15, owner{"Hiren Vaidya"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerinfo.name)

}
