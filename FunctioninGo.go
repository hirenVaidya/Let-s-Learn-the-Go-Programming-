package main

import "fmt"

type user struct {
	Name string
	Age  int
}

// Pointer receiver: can modify the struct
func (u *user) Birthday() { u.Age++ }

// Value receiver: works on a copy
func (u user) Greet() string { return "Hi, " + u.Name }

func main() {
	u := user{Name: "Hiren", Age: 25}
	u.Birthday()
	fmt.Println(u.Greet(), u.Age) // Hi, Ravi 26

	users := []user{u, {Name: "Priya", Age: 30}} // slice
	users = append(users, user{"Amit", 22})

	byName := map[string]user{} // map
	for _, x := range users {
		byName[x.Name] = x
	}
	if v, ok := byName["Priya"]; ok { // "comma ok" idiom
		fmt.Println(v.Age)
	}
}
