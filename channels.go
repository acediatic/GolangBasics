package main

import (
	"fmt"
	"sync"
)

var wg3 = sync.WaitGroup{}

func channels() {
	ch := make(chan int)
	for j := 0; j < 10; j++ {
		wg3.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg3.Done()
		}()
		go func() {
			ch <- 42 // pauses till space available in the channel (unbuffered channel)
			wg3.Done()
		}()
		wg3.Wait()
	}
}

// receive / send only
func channels2() {
	ch := make(chan int)
	wg3.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg3.Done()
	}(ch)
	go func(cd chan<- int) {
		ch <- 42 // pauses till space available in the channel (unbuffered channel)
		wg3.Done()
	}(ch)
	wg3.Wait()
}

// Buffering

func buffering() {
	ch := make(chan int, 2)
	wg3.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}

		// alternative
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				// channel closed
				break
			}
		}
		wg3.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 43
		close(ch)
		wg3.Done()
	}(ch)
	wg3.Wait()
}

// select statements

type logEntry struct {
	msg string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // empty struct. Allocates no memory, can only see if message sent
// hence used as a signal channel

func selectMain() {
	go logger()
	logCh <- logEntry{"Hello"}
	logCh <- logEntry{"World"}
	doneCh <- struct{}{} // pass in empty struct, to signal end
}

func logger() {
	for { // infinite loop
		select { // blocking, unless you have a default case
		case entry := <-logCh:
			fmt.Println(entry.msg)
		case <-doneCh:
			break // if message, end loop
		}
	}
}
