package main

import "fmt"

type Vehicle struct {
	color, model, company string
}

func main() {
	// m = make(map[int]Vehicle)

	/* Need to define the map, as nil map can not be assigned any value*/
	var m map[string]Vehicle

	/* create a map*/
	/// Syntax 1
	// m = make(map[string]Vehicle)
	// m["USA"] = Vehicle{"Red", "GT", "Ford"}
	// m["India"] = Vehicle{"Red", "Bolero", "Mahindra"}
	// m["Germany"] = Vehicle{"White", "Polo", "Volkswagen"}

	/// Syntax 2
	m = map[string]Vehicle{
		"India":   {"Red", "Bolero", "Mahindra"},
		"USA":     {"Red", "GT", "Ford"},
		"Germany": {"White", "Polo", "Volkswagen"},
	}

	fmt.Println(m)

	// Get values from a map
	vehicle, ok := m["India"]
	fmt.Println(vehicle)
	fmt.Println(ok)

	// Delete a key from map
	delete(m, "India")
	fmt.Println(m)

	var str string = "Hello"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

}
