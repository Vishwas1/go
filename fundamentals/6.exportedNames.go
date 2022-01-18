package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

/// Not exported outside this package
func notExportedFunction() {

}

//// Exported out side this package
// If you want functions to be exported outside,
// make sure to use first letter as caps letter of the name of the function
func ExportedFunction() {

}
