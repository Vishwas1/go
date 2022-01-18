package main

import (
	"fmt"
	"time"
)

func main() {
	num := 1

	// if-else
	if num > 3 {
		fmt.Println("Hi")
	} else {
		fmt.Println("Bye")
	}

	// Switch
	// Notice break statement is already provided
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("Day after tomorrow")
	default:
		fmt.Println("To far")
	}

	// Defer statement
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer fmt.Println("Defer stmt: This will be printed later like in the end")
	fmt.Println("This will be printed first")
	fmt.Println("This will be printed second")
	fmt.Println("This will be printed third")

	// Deferred function calls are pushed onto a stack.
	// When a function returns, its deferred calls are executed in last-in-first-out order.
	defer fmt.Println("Defer stmt: This will be second deferedd statement")
	defer fmt.Println("Defer stmt: This will be first deferedd statement")

}
