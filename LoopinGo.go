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

	package main

import "fmt"

func main() {
    // Basic switch: no break needed, cases don't fall through
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Other day")
    }

    // Multiple values per case
    switch day {
    case 6, 7:
        fmt.Println("Weekend")
    case 1, 2, 3, 4, 5:
        fmt.Println("Weekday")
    }

    // Switch with no expression (cleaner than if-else chains)
    score := 82
    switch {
    case score >= 90:
        fmt.Println("A")
    case score >= 80:
        fmt.Println("B")
    case score >= 70:
        fmt.Println("C")
    default:
        fmt.Println("F")
    }

    // Switch with init statement
    switch os := "linux"; os {
    case "darwin":
        fmt.Println("macOS")
    case "linux":
        fmt.Println("Linux")
    default:
        fmt.Println("Other")
    }

    // fallthrough: forces execution of the next case (rarely used)
    switch 1 {
    case 1:
        fmt.Println("one")
        fallthrough
    case 2:
        fmt.Println("two (via fallthrough)")
    case 3:
        fmt.Println("three") // not printed
    }

    // Type switch (you'll use this with interfaces later)
    var v any = 3.14
    switch t := v.(type) {
    case int:
        fmt.Println("int", t)
    case string:
        fmt.Println("string", t)
    case float64:
        fmt.Println("float64", t)
    }
}
}
