package main

import "fmt"

const MAX int = 3

func main() {

	fmt.Println("======== Basic concept of pointers ==============")
	var a int = 10
	fmt.Printf("Value  of a variable: %v\n", a)

	//// Address of memory location of a variable
	// every variable is a memory location and every memory location has its address defined
	// which can be accessed using ampersand (&) operator, which denotes an address in memory.
	// prints the memory address of variable a
	fmt.Printf("Address of a variable: %x\n", &a)

	//// Pointer
	// A pointer is a variable whose value is the address of another variable
	var ap *int = &a /* store address of a in pointer variable*/
	// ap := &a //  the above statement can also be written in this way

	fmt.Printf("Address stored in ap variable: %x\n", ap)

	fmt.Printf("Value of *ap variable: %d\n", *ap)

	/// Nil Pointer
	var p *int
	fmt.Println(p)

	fmt.Println("======== Array of pointers ==============")
	arrayPointers()

	fmt.Println("======== Call by Reference ==============")
	var var1 int = 100
	var var2 int = 200

	fmt.Printf("Before swap, value of var1 : %d\n", var1)
	fmt.Printf("Before swap, value of var2 : %d\n", var2)

	// Call by refrence
	swap(&var1, &var2)

	fmt.Printf("After swap, value of var1 : %d\n", var1)
	fmt.Printf("After swap, value of var2 : %d\n", var2)

	fmt.Println("======== Pointers to Pointers ==============")
	pointerToPointer()
}

/// Array of pointers
func arrayPointers() {
	a := []int{10, 100, 200}
	for i := 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, a[i])
	}

	var ptrA [MAX]*int
	for j := 0; j < MAX; j++ {
		ptrA[j] = &a[j]
	}

	for k := 0; k < MAX; k++ {
		fmt.Printf("Value of a[%d] = %d and address of a[%d] = %d\n", k, *ptrA[k], k, ptrA[k])
	}

}

/// Passing pointers to a function
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

/// Pointer to pointer
func pointerToPointer() {

	// Normally a pointer contains the address of a variable
	// In pointer to pointer - first pointer contains the address of a second pointer
	// and the second pointer contains the address of the acual variable

	a := 10

	var ptrA *int
	var pptrA **int

	ptrA = &a
	pptrA = &ptrA

	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *ptrA = %d\n", *ptrA)
	fmt.Printf("Value available at **pptrA = %d\n", **pptrA)

}
