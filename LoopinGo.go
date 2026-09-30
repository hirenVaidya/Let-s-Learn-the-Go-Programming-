package main

import "fmt"

func main() {
	// Form 1: classic (init; condition; post)
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// Form 2: like a while loop (condition only)
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("n =", n) // 128

	// // Form 3: infinite loop (use break to exit)
	count := 0
	for {
		count++
		if count == 3 {
			continue // skip the rest of this iteration
		}
		if count > 5 {
			break // exit the loop
		}
		fmt.Println("count =", count)
	}

	// Form 4: range over a slice
	fruits := []string{"apple", "banana", "mango"}
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}

	// Ignore index with _
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	frutis := []string{"apple", "banana", "mango"}
	for _, something := range frutis {
		fmt.Println(something)
	}
	// Range over a string (gives runes, not bytes)
	for i, ch := range "Go!" {
		fmt.Printf("%d: %c\n", i, ch)
	}

	// Range over a map (order is random)
	ages := map[string]int{"Ravi": 25, "Priya": 30}
	for name, age := range ages {
		fmt.Println(name, age)
	}

	// Range over an integer (Go 1.22+)
	for i := range 3 {
		fmt.Println("round", i) // 0, 1, 2
	}
}
