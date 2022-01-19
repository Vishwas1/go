package main

import "fmt"

//// A struct is a collection of fields. Think of it as defining cutom type
/// Syntax 1
type Vehicle struct {
	color   string
	model   string
	company string
}

/// Syntax 2
// type Vehicle struct {
// 	color, model, company string
// }

func main() {

	// declaring variable of the type Vehicle
	v := Vehicle{"Red", "GT", "Fiat"}
	fmt.Println("I have", v.company, v.model, "in", v.color, "color")

	/// pointer to struct
	var p *Vehicle
	p = &v
	// p := &v /* you can either write above two lines or just us this syntax*/
	// getting values
	fmt.Println("I have", (*p).company, (*p).model, "in", (*p).color, "color")
	// you can use without dereferencing
	fmt.Println("I have", p.company, p.model, "in", p.color, "color")

	t := Vehicle{}
	fmt.Println(t)

	w := Vehicle{company: "Ford", model: "GT"}
	fmt.Println(w)

	// Array of structs
	var vehicles = []Vehicle{
		{"Red", "GT", "Ford"},
		{"Black", "Scorpio", "Mahindra"},
		{"Green", "GT", "Fiat"},
	}

	for i := 0; i < len(vehicles); i++ {
		fmt.Println("I have", vehicles[i].company, vehicles[i].model, "in", vehicles[i].color, "color")
	}
}
