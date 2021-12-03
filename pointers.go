package main

import (
	"fmt"
)

// asterix (*) before type means it's a pointer to that type
// asterix before a variable means it's the dereference of that variable (get value at that address)
// The 0 value of a pointer is <nil>

func pointers() {
	var a int = 100
	var b *int = &a   // b is a pointer to the memory address where a is stored
	fmt.Println(a, b) // 100, 0xsomeHexString (memory address of a)
	a = 26
	fmt.Println(a, b) // 26, 26
}

// Does not have pointer arithmetic as it leads to unsafe code. If really need, use the unsafe
// package.

func pointersAndStructs() {
	var ms *myStruct
	ms = new(myStruct)         // option 1: nill
	ms = &myStruct{foo: "bar"} // option 2: create object and pass the address to it
	(*ms).foo = "bar"          // derefernce operator has lower precedence than the dot operator.
	ms.foo = "bar2"            //  The compiler AUTOMATICALLY DEREFERNCES THE POINTER
	fmt.Println(ms)

}

type myStruct struct {
	foo string
}
