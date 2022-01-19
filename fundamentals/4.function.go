package go/fundamentals

import (
	"fmt"
	"math"
)

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

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
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

	// Function used as variable
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// Funcation passed as parameter to another funcation
	fmt.Println(compute(hypot))

}
