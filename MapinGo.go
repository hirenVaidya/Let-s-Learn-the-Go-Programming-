package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["Hiren"] = 25
	ages["Alka"] = 19

	fmt.Println(" alka age is just : " + fmt.Sprint(ages["Alka"]))

}
