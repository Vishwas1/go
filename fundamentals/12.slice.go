package main

import "fmt"

func main() {

	//// An array has a fixed size. A slice, on the other hand,
	//// is a dynamically-sized, flexible view into the elements of an array.
	//// Go Slice is an abstraction over Go Array. Go Array allows you to define variables that can hold several data items of the same kind but it does not provide any inbuilt method to increase its size dynamically or get a sub-array of its own. Slices overcome this limitation. It provides many utility functions required on Array and is widely used in Go programming.

	// This is an array (of fixed size)
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Slice
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// This selects a half-open range which includes the first element, but excludes the last one.
	var s []int = primes[1:4]
	// prints [3 5 7]
	fmt.Println(s)

	weekdays := [5]string{"Monday", "Tuesday", "Webnesday", "Thrusday", "Friday"}
	// prints [Tuesday Webnesday Thrusday]
	s1 := weekdays[1:4]
	fmt.Println(s1)

	//// A slice does not store any data,
	// it just describes a section of an underlying array.
	// so if we change anything in slice it will reflect in array too
	s1[1] = "Sunday"
	// prints [Tuesday Sunday Thrusday]
	fmt.Println(s1)
	// prints [Monday Tuesday Sunday Thrusday Friday]
	fmt.Println(weekdays)

	/// Slice literals
	// are like normal array without size
	weekends := []string{"Saturday", "Sunday"}
	fmt.Println(weekends)

	/// Slice has length and capacity
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array,
	// counting from the first element in the slice.
	// prints len=3 cap=4 slice=[Tuesday Sunday Thrusday]
	// notice cap is 4 and not 5 since the slice starts from 1 . see line 23

	fmt.Printf("len=%d cap=%d slice=%v\n", len(s1), cap(s1), s1)

	fmt.Println("========= Make slice =============")
	makeSlice()

}

func makeSlice() {
	//// Slices can be created with the built-in make function;
	//// The make function allocates a zeroed array and returns a slice that refers to that array:

	a := make([]int, 5)
	// prints a len=5 cap=5 [0 0 0 0 0]
	printSlice("a", a)

	b := make([]int, 0, 5)
	// prints b len=0 cap=5 []
	printSlice("b", b)

	c := b[:2]
	// prints c len=2 cap=5 [0 0]
	printSlice("c", c)

	d := c[2:5]
	// prints d len=3 cap=3 [0 0 0]
	printSlice("d", d)

	fmt.Println("========= Slice Append =============")
	//// Appending to slice
	var s []int
	printSlice2(s)

	s = append(s, 12)
	printSlice2(s)

	s = append(s, 20)
	s = append(s, 24)
	s = append(s, 30)
	s = append(s, 60, 40, 50) /*Adding more than one elm*/
	printSlice2(s)

	fmt.Println("========= Slice Range =============")
	/// Range
	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a copy of the element at that index.
	// Here i is the index and v is the value at that index
	for i, v := range s {
		fmt.Println("Slice item", i, "is", v)
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
