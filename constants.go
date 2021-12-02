package main

import (
	"fmt"
)

// shadowing also applies, but not recommended (effectively changing constant value)
const myConst int = 10

func constants() {
	// constants naming convention same as variable naming convention
	const myConst int = 10

	// myConst = 20 // error: cannot assign to myConst

	const a = 42
	var b int16 = 27

	// can do this because the type is inferred to be int16 for constant
	// different to what you can do with variables
	// (which have a default value for int of at least 32)
	fmt.Printf("%v, %T\n", a+b, a+b)
}

// enums
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
	d
) // would assume error, but compiler infers, and makes d = 3

// starts again at 0
const (
	a2 = iota // 0
	b2 = iota // 1
)

func enums() {
	var myEnum int = 1
	var myErrorEnum int

	// acts as expected
	fmt.Printf("%v, %T\n", myEnum, myEnum)

	// will print out the value 0, even though it wasn't initialized
	fmt.Printf("%v, %T\n", myErrorEnum, myErrorEnum)
}

// fix by making the first value to a error
const (
	errorEnum = iota
	enum1
	enum2
)

// or use a write-only first value
const (
	_ = iota
	enum11
	enum12
)

// can set the first value of iota.
const (
	_ = iota + 5
	enum21
	enum22
)

// can also use bit shifting (cause can't use math package)
const (
	_  = iota // ignore first value
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func PrintFileSize(size int) {
	switch {
	case size < KB:
		fmt.Printf("%d B\n", size)
	case size < MB:
		fmt.Printf("%d KB\n", size/KB)
	case size < GB:
		fmt.Printf("%d MB\n", size/MB)
	case size < TB:
		fmt.Printf("%d GB\n", size/GB)
	case size < PB:
		fmt.Printf("%d TB\n", size/TB)
	}
}

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func Flags() {
	var roles byte = isAdmin | isHeadquarters | canSeeFinancials
	fmt.Printf("%b\n", roles)
	// 100101

	fmt.Printf("Is Admin? %v\n", roles&isAdmin == isAdmin)        //true
	fmt.Printf("Is HQ? %v\n", roles&canSeeAfrica == canSeeAfrica) //false
}
