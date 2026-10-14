package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type contractInfo struct {
	ID   string
	Name string
}

type Purchase struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	var contracts []contractInfo = loadJSON[contractInfo]("./contracts.json")
	fmt.Println("contracts:", contracts)

	var purchases []Purchase = loadJSON[Purchase]("./purchases.json")
	fmt.Println("purchases:", purchases)
}

func loadJSON[T contractInfo | Purchase](filename string) []T {
	data, _ := ioutil.ReadFile(filename)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)
	return loaded

}
