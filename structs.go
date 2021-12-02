package main

import (
	"fmt"
)

// struct is pass by value

// Doctor is a struct
/*
	in Go, a struct is a collection of fields
	each field has a name and a type
	the type can be a primitive type, a struct, or a pointer to a struct
	the name is the name of the field
*/

// same convention on variable names
type Doctor struct {
	name   string
	age    int
	number int
}

func structs() {
	aDoctor := Doctor{
		name:   "Dr. Strange",
		age:    42,
		number: 123456789,
	}
	fmt.Println(aDoctor)

	// accessing fields
	fmt.Println(aDoctor.name)
}

// annoymous struct, usually for only very short lived structs -> times when you don't need a formal
// type available throughout the package.
func annoymousStruct() {
	aDoctor := struct{ name string }{name: "Dr. Strange"}
	fmt.Println(aDoctor)
}
