package go/fundamentals

import "fmt"

// Package level scope
var result int

//// Syntax 1
func add(x int, y int) int {

	// Function level scope
	var res int = x + y
	result = res

	return res
}

func main() {

	num1 := 5
	num2 := 4

	// prints 9
	fmt.Println(add(num1, num2))

	// prints 9
	fmt.Println(result)
}
