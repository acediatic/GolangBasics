package main

import (
	"fmt"
)

func loops() {
	// increment operator is a statement, not an expression
	// so can't use multiple in one statement.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// can manipulate counter in the loop, but don't. Not a good idea.

	// can skip initializing counter and incrementing i nthe while loop.
	// i is now scoped to the loops function
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// how go does the while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	for {
		fmt.Println("infinite loop")

		// break statement
		break

		// continue statement
		continue
	}

	// use loop label to break out of (or continue) nested loops
outerLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break outerLoop
			}
			fmt.Println(i, j)
		}
	}

	collection := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for in loop
	// always get key and value
	for k, v := range collection {
		fmt.Println(k, v)
	}

	for k := range collection {
		fmt.Println(k)
	}

	for _, v := range collection {
		fmt.Println(v)
	}

}
