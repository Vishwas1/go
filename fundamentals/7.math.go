package go/fundamentals

import (
	"fmt"
	"math"
)

func main() {
	var x float64 = 26
	result := math.Sqrt(x)
	fmt.Println("Square root of", x, "=", result)

	// printing till 2 decimal places
	fmt.Printf("Square root is %.2f", result)

	fmt.Println()

	// floor
	fmt.Println("Square root of", x, "in flor=", math.Floor(math.Sqrt(x)))

	fmt.Println("Square root of", x, "=", math.Ceil(math.Sqrt(x)))

	fmt.Println("Square root of", x, "=", math.Round(math.Sqrt(x)))
}
