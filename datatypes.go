package main

import (
	"fmt"
)

func Datatypes() {
	// boolean
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	var uninitalizedBool bool
	fmt.Printf("%v, %T\n", uninitalizedBool, uninitalizedBool)

}

func TypeConversion() {
	// integer is always _at least_ 32 bits
	// int8
	// int16
	// int32
	// int64

	// equivalent unsigned integer for signed integer
	// uint8

	// Integer division produces integer (truncated towards 0)

	// must type convert in _all_ cases:

	var a int = 10
	var b int8 = 3

	// Error:
	// fmt.println(a + b)

	// Fine:
	fmt.Println(a + int(b))
}

func BinaryOps() {
	a := 10 // 1010
	b := 3  // 0011
	c := 8  // 1000 , 2^3

	// AND
	fmt.Println(a & b) // 0010
	// OR
	fmt.Println(a | b) // 1011
	// XOR
	fmt.Println(a ^ b) // 1001
	// AND NOT (only if neither bit set)
	fmt.Println(a &^ b) // 0100
	// Bit Shift left
	fmt.Println(c << 2) // 2^3 * 2^2 = 2^5 = 32 = 100000
	// Bit Shift right
	fmt.Println(c >> 2) // 2^3 / 2^2 = 2^1 = 2 = 10
}

func floatingPoint() {
	// float32
	// float64

	n := 3.14
	n = 13.7e12
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)
}

func strings() {
	// strings in go are just aliases for bytes.
	// strings are immutable
	// they are declared using double quotes (")
	// they are UTF-8 encoded

	s := "this is a string"

	// gives 105, uint8
	fmt.Printf("%v, %T\n", s[2], s[2])

	// gives i, uint8
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// cannot edit strings, they're immutable
	// can however concatenate strings
	// or convert to slice of bytes which can then be edited

	b := []byte(s)
	// collection of byte slices, of type []uint8
	fmt.Printf("%v, %T\n", b, b)
}

func runes() {
	// runes are characters
	// they are UTF-32 encoded
	// they are stored as int32
	// they are declared using single quotes (')

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
	// gives 92, int32
	// runes are just an int32

	// can use #Reader.ReadRune() to read runes
}
