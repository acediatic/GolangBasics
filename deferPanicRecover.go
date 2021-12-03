package main

import (
	"fmt"
)

/*
	Defer is used to ensure that a function call is performed later in a program’s execution.
	Deferred statements are executed after the final return statement of a function, before it returns,
	and from there in Last In First Out order.

	They execute AFTER the function is done, but BEFORE it returns to the calling function.

	In practice, this allows use to associate the opening of a resource and the closing of
	a resource right next to each other. This improves readability, and reduces bugs in code.

	Defer takes a FUNCTION CALL, not a function definition.
*/
func deferFn() {
	fmt.Println("start")        // 1st
	defer fmt.Println("middle") // 3rd
	fmt.Println("end")          // 2nd
}

/*
	Deferred functions take the values from the time they were called
*/
func deferVars() {
	a := "start"
	defer fmt.Println(a) // prints "start"
	a = "end"
}

// Panic - equivalent of an exception, but go uses less of them (instead favouring returning error codes)
// panics happen after defer statements are executed
func panicing() {
	// use built in panic function
	fmt.Println("start")
	defer fmt.Println("middle")
	panic("PANIC")
	fmt.Println("end")

	/*
		console:
		start
		middle
		panic: PANIC
	*/

}

/*
	Recover - recover from a panic by calling recover()

	In order to work out what the error was, we need to call recover().
	This means were committing to dealing with it. If we can't, we need
	to repanic the application by calling panic again
*/
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from", r)
		}
	}()
}
