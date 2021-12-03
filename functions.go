package main

import (
	"fmt"
)

// Can pass in/around pointers too

func withParams(varName int) {
	fmt.Println(varName)
}

// list of variables with the type declared last, assumes same type for all
func allSameType(a, b, c string) {
	fmt.Println(a, b, c)
	fmt.Println("All same type")
}

func callVarArgs() {
	variableArgs("hello", 1, 2, 3, 4)
}

// can only have one variadic parameter, has to be last parameter
func variableArgs(msg string, values ...int) {
	// values is a slice that holds all the passed in values
	for _, value := range values {
		fmt.Println(value)
	}
}

// return value type goes after parenthesis before curly brace
func returnValue() (string, int) {
	return "Hello", 1
}

// named return values
func namedReturn() (msg string) {
	msg = "Hello"
	return // returns the value msg, but may be less readable.
}

/*
	in many programs, this would be dangerous, as you're returning a pointer to a memory address
	on the stack. However, Go automatically promotes the address to the heap when it's returned
	by the function
*/
func returningPointer() *int {
	// return a pointer
	var a int = 1
	return &a
}

// process error value in calling function
func returningErrorValue(a, b float64) (float64, error) {
	if a != 0.0 && b != 0.0 {
		return a / b, nil // explicity return nil if nothing went wrong
	} else {
		return 0, fmt.Errorf("Error") // return error if something went wrong
	}
	// do error checknig as soon as possible, and return out as soon as possible (left justify)
}

// annoymous function
func anon() {
	func() {
		fmt.Println("Hello")
	}()
}

func funcAsVar() {
	// function literal
	var f func() = func() {
		fmt.Println("Hello")
	}
	f()
}

// *** Method calls ***
type Person struct {
	Name string
}

// Method gets passed context (simiar to self in Python)
// It's passing by value, not by pointer
// Can also pass in pointer receiver
func (p Person) SayHello() {
	fmt.Println("Hello", p.Name)
}

// pointer receiver
func (p *Person) SayHello2() {
	fmt.Println("Hello", p.Name)
}

func greetPerson() {
	p := Person{Name: "John"}
	p.SayHello()
}
