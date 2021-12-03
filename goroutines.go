package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Go routine: lightweight thread abstraction, cheap to create/destroy
// Go automatically maps "Go Routines" to the underlying OS threads.
// This avoids the overhead of thread creation.

// Can check for race condition issues at compile time by running
// go run -race myFile.go

func parallelmain() {
	runtime.GOMAXPROCS(100) // sets the number of OS threads to 100
	var msg = "hello"
	go func() {
		println(msg) // Creates race condition, often prints "World"
	}()
	msg = "world" // because we change the message here
	time.Sleep(100 * time.Millisecond)
}

func fixedParallel() {
	var msg = "hello"
	go func(msg string) {
		fmt.Println(msg) // passes value, no race condition
	}(msg)
}

// Wait Groups: designed to synchornise go routines
// Essentially a semaphore

// wait group increment/decrement is threadsafe
var wg = sync.WaitGroup{}

func waitingFn() {
	var msg = "hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() // decrements the wait group counter
	}(msg)
	wg.Wait() // blocks until the wait group counter is 0
}

// Mutexs also exist
var m = sync.RWMutex{} // infinite readers, only one writer.

var wg2 = sync.WaitGroup{}
var counter = 0

func main2() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 100; i++ {
		wg2.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg2.Wait()
}

func sayHello() {
	fmt.Println("Hello")
	m.RUnlock()
	wg2.Done()
}

func increment() {
	m.Lock()
	counter++
	m.Unlock()
	wg2.Done()
}
