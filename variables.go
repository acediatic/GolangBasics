package main

import (
	"fmt"
	"strconv"
)

// Have to use full declaration syntax
var pacakgeLevelVariable string = "package level variable"

// Three scopes: package, function, and block

var I int = 10 // exposed outside the package
var i int = 10 // scoped to the package

// variable block
var (
	actorName  string = "Tom Hanks"
	actorAge   int    = 56
	actorMovie string = "Forrest Gump"
)

func main() {
	// local declaration
	// variable, name, type
	// due to shadowing, local variable takes precedence over global
	var i int

	// initialization
	i = 42

	// declaration and initialization
	// easier to read, but type is inferred
	k := 99 //. adding a . afterwards makes it of type int64 rather than 32-bit int
	// l := 99.

	var j float32
	j = float32(i) // type conversion

	var x string
	x = strconv.Itoa(i) // type conversion

	fmt.Println(i)
	fmt.Println(k)
	fmt.Println(x)
	fmt.Println(j)
}
