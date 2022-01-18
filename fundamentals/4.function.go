package main

import "fmt"

//// Syntax 1
func add(x int, y int) int {
	return (x + y)
}

//// Syntax 2
func add2(x, y int) int {
	return x + y
}

//// Syntax 3:  Returning more than one values
func addSub(x, y int) (int, int) {
	return x + y, x - y
}

//// Syntax 4:  Naked return;  more than one values
func addSub2(x, y int) (res1 int, res2 int) {
	res1 = x + y
	res2 = x - y
	return
}

func main() {

	num1 := 5
	num2 := 4

	// prints 9
	fmt.Println(add(num1, num2))
	// prints 9
	fmt.Println(add2(num1, num2))

	addition, subtraction := addSub(num1, num2)
	// prints 9 1
	fmt.Println(addition, subtraction)

	addition1, subtraction1 := addSub2(num1, num2)
	// prints 9 1
	fmt.Println(addition1, subtraction1)

}
