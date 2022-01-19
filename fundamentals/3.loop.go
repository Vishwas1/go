package go/fundamentals

import (
	"fmt"
)

func main() {
	//// Syntax 1: Looping infinite times
	// for {
	// 	fmt.Println("Printing multiple times")
	// }

	//// Syntax 2: Looping finite times
	fmt.Println("Using loop in lengthy way...")
	i := 1
	for i <= 5 {
		fmt.Println("Printing i = ", i)
		i++
	}

	//// Syntax 3: Looping finite times, more compact way
	fmt.Println("Using loop in compact way...")
	for i := 1; i <= 5; i++ {
		fmt.Println("Printing i = ", i)
	}

}
