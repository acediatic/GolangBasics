package main

import (
	"fmt"
)

// PREFERENCES
/*
	1. Use many smaller rather than fewer larger interfaces
	2. Don't export interfaces for types that will be consumed (unless you want to, use concrete type)
	3. Export interfaces for types that will be used by a package
	4. Design functions and methods to receive interfaces whenever possible (SRP)
*/

/*
	Naming convention: If have a one method interface, name it the method name with 'er' suffix.
	Such as write -> Writer
*/

func mainFn() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello, World!"))
}

type Writer interface {
	Write([]byte) (int, error)
}

// don't explicity implement interfaces, instead implicitly
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(p []byte) (int, error) {
	n, err := fmt.Print(string(p))
	return n, err
}

// dont have to use structs

func runIncrementer() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// Can embed interfaces in other interfaces
// then you implement the enclosed interface, and can use it accordingly

// empty interface
// interface{} an interface defined on the fly that has no methods
// can be used to represent any type
// it doesn't have any use methods however, so we can't do anything with it.
// almost always an intermediate step

// type switch
func typeSwitch() {
	var i interface{} = "hello"
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

/*
NOTE: when implementing an interface, if you use a concrete value to implement the interface,
it will have to all have VALUE receivers (can't handle pointers).

If you use a pointer to a concrete value, it will have POINTER & VALUE receivers.
The method set of a value reciever is all the methods that have a value type AND all the methods
that have a pointer type.

i.e. 	1) var wc WriterCloser = myWriterCloser{}
		2) var wc WriterCloser = &myWriterCloser{}
*/
