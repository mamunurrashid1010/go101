package main

import (
	"fmt"
	"strings"
)

// ============================================================
// TOPIC: Go Basics - Variables, Types, Control Flow, Strings
// ============================================================

// --- EXAMPLE ---

var course = "Go 101" // package-level variable
const maxRetries = 5  // constant

func main() {
	fmt.Println("Welcome to", course)

	// Variables & short declaration
	name := "Ada"
	age := 25
	isStudent := true
	fmt.Printf("name=%s age=%d student=%t\n", name, age, isStudent)

	// String formatting
	greeting := fmt.Sprintf("Hello, %s!", strings.ToUpper(name))
	fmt.Println(greeting)

	// Multiple return values + error handling
	result, err := add(21, 21)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	fmt.Println("21 + 21 =", result)

	// Switch statement (no break needed in Go)
	switch {
	case result > 50:
		fmt.Println("Big number!")
	case result == 42:
		fmt.Println("The answer to everything!")
	default:
		fmt.Println("Just a number:", result)
	}

	// For loop + defer (LIFO order)
	for i := 0; i < 3; i++ {
		defer fmt.Println("deferred:", i)
		fmt.Println("loop iteration:", i)
	}

	// Practice question
	fmt.Println("PRACTICE QUESTIONS: ")

	// Q1:
	fmt.Println("maxRetries = ", maxRetries)

	//Q2:
	fmt.Println("6 is Even: ", isEven(6))
	fmt.Println("3 is Even: ", isEven(3))
	fmt.Println("10 is Even: ", isEven(10))

	// Q3:
	fmt.Println(greet("Mamunur Rashid"))

	// Q4:
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd number: ", i)
	}

	// Q5:
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Printf("%s is a weekday.\n", day)
	case "Saturday", "Sunday":
		fmt.Printf("%s is a weekend.\n", day)
	default:
		fmt.Println("Unknown day")
	}

	// Q6:
	/*
		Output: 2, 1, 0
		Each defer pushes the call onto a stack with the current value of i captured as an argument.
		When main() returns, the stack unwinds from the top, so the most recently deferred call (i=2) runs first.
	*/

}

func add(a, b int) (int, error) {
	sum := a + b
	if sum < 0 {
		return 0, fmt.Errorf("negative sum: %d", sum)
	}
	return sum, nil
}

func isEven(n int) bool {
	return n%2 == 0
}

func greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go 101.", name)
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Declare a constant called `maxRetries` with value 5.
//     Print it inside main().
//
// Q2: Write a function `isEven(n int) bool` that returns true
//     if n is even. Call it from main() with a few test values.
//
// Q3: Write a function `greet(name string) string` that returns
//     "Hello, <name>! Welcome to Go 101." using fmt.Sprintf.
//
// Q4: Use a for loop to print numbers 1 to 10. Skip even
//     numbers using `continue`.
//
// Q5: Write a switch statement that takes a day string
//     ("Monday", "Saturday", etc.) and prints whether it's
//     a weekday or weekend.
//
// Q6: What will this print? Why?
//     for i := 0; i < 3; i++ {
//         defer fmt.Println(i)
//     }
//     (Answer: 2, 1, 0 - defer is LIFO)
//
// ============================================================
