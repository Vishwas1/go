package go/fundamentals

import "fmt"

func main() {
	//// syntax 1
	var i, j int

	// this will print 0
	fmt.Println(i + j)

	//// syntax 2
	var a, b int = 5, 10

	// this will print 15
	fmt.Println(a + b)

	//// syntax 3:  same as var x = 10
	x := 10

	// this will print 25
	fmt.Println(a + b + i + j + x)

	/// Const
	const ENV = "YOU CAN NOT CHANGE THIS"
	fmt.Println(ENV)

	var days = []string{"hello", "hi", "test", "new", "old"}
	printDays(days)

}

func printDays(days []string) {
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
}
