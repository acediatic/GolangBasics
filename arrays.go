package main

import "fmt"

func arrays() {
	grades := [3]int{97, 85, 93}

	// if initalising array literal, you can omit the size
	// creates array with enough size to hold all values
	grades2 := [...]int{97, 85, 93}
	grades2[1] = 100

	// or create an empty array
	var students [3]string
	students[0] = "John"

	// calculate the size of an array
	fmt.Println(len(students))

	// can also create 2d arrays (matrix)
	// 2d arrays are arrays of arrays
	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(identityMatrix)

	// in Go, arrays are considered data.
	a := [3]int{1, 2, 3}
	b := a // creates a complete copy of the array
	b[1] = 4
	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [1 4 3]

	// can choose to pass pointer using &
	c := &a
	c[1] = 5
	fmt.Println(a) // [1 5 3]
	fmt.Println(c) // [1 5 3]

	fmt.Printf("%v\n", grades)
}

func slice() {
	// slices are like arrays but they can be resized
	a := []int{1, 2, 3, 4, 5}

	// slicing operations can have as their source an array or a slice
	// or the make function
	// make([]int, len, cap)
	makeTest := make([]int, 5, 10)
	fmt.Println(makeTest)

	// also get a function called capacity. This determines the current size of the slice
	fmt.Println(cap(a)) // 5

	// slices are REFERENCE types, so are pointing to the same underlying array

	e := a[:]   // [1 2 3 4 5] All elements
	b := a[1:3] // [2 3 4] second, third, and fourth elements
	c := a[:3]  // [1 2 3] first three elements
	d := a[1:]  // [2 3 4 5] second and following elements

	fmt.Println(e, b, c, d)

	// functions on slices

	// append adds an element to the end of the slice
	// append(slice, element)
	// creates a new array if the capacity is too small
	// does so by doubling in powers of two
	a = append(a, 1) // [1 2 3 4 5 1]

	// can also use spread operator to spread out values of the array/slice
	a = append(a, []int{2, 3, 4, 5}...)

	// removing value from the middle of a slice
	a = append(a[:2], a[3:]...)

	// IF REMOVING VALUES FROM ARRAY, check there are no references. Otherwise can get weird results:
	test := []int{1, 2, 3, 4, 5}
	fmt.Println(test) // [1 2 3 4 5]
	test2 := append(test[:2], test[3:]...)
	fmt.Println(test2) // [1 2 4 5]
	fmt.Println(test)  // [1 2 4 5 5]
}
