// %d	integer
// %s	string
// %f / %.2f	float
// %t	boolean
// %v	any value (default)
// %+v	struct with field names
// %T	type

package main

import "fmt"

func main() {
	v := "There is no place like 127.0.0.1"
	fmt.Println(v)

	name := "Hiren(Hero)"
	age := 25

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Value: %v\n", age)
	fmt.Printf("Speed of Light : %.2f\n", 299792458.0)
	fmt.Printf("Quoted: %q\n", name)

	msg := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(msg)

}
