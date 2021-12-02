package main

import (
	"fmt"
	"math"
)

func if_and_switch() {
	//if
	// not allowed to have a single line block, need curly braces
	if true {
		fmt.Println("if")
	}

	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
	}

	// pop is scoped to the block
	if pop, ok := statePopulations["California"]; ok {
		fmt.Println(pop)
	} else if true {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	// operators
	// > < >= <= == !=
	// || &&

	// floating point arithmetic can cause unexpected results, with exact equality.
	myNum := 0.123
	if math.Abs((myNum/math.Pow(math.Sqrt(myNum), 2))-1) < 0.00001 {
		fmt.Println("same number")
	} else {
		fmt.Println("different number")
	}

	// switch
	// with explicit definition, cases can't overlap
	// break keyword is implicit (to avoid issues with fallthrough)
	switch 2 {
	case 1, 5, 10:
		fmt.Println("one")
	case 2, 3, 8: // replace with 5 to see can't overlap behaviour
		fmt.Println("two")
	default:
		fmt.Println("default")
	}

	// can use initialiser
	switch myNum := 2; myNum {
	case 1, 5, 10:
		fmt.Println("one")
	default:
		fmt.Println("default")
	}

	// switch with no condition
	// conditions can overlap
	// use keyword fallthrough to allow fallthrough
	// but not usually used considering you can have mutliple conditions.
	// use the break keyword to leave the switch statement early
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than 10")
		fallthrough
	case i <= 20:
		fmt.Println("greater than 10")
		break
		fmt.Println("this should not print")
	default:
		fmt.Println("equal to 10")
	}

	// switching by type
	var t interface{} = 1 // type integer
	switch t.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case [3]int:
		fmt.Println("3 array")
	case [4]int:
		fmt.Println("4 array") // different to above 3 int
	default:
		fmt.Println("unknown")
	}

}
