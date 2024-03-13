// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing(c chan int) { // specify the type of the channel
	//TODO: increment i 1000000 times
	for j := 0; j < 999999; j++ {
		c <- 1
	}
}

func decrementing(c chan int) { // specify the type of the channel
	//TODO: decrement i 1000000 times
	for j := 0; j < 1000000; j++ {
		c <- 0
	}
}

func omnicrementing(myChan chan int) {
	for {
		select {
		case inc := <-myChan:
			if inc == 1 {
				i++
			} else {
				i--
			}
		}
	}
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1? Set the amount of threads that can run go-routines simultanoiusly
	runtime.GOMAXPROCS(2)

	myChan := make(chan int)

	// Spawn both functions as goroutines
	go incrementing(myChan) // pass the channel as an argument
	go decrementing(myChan) // pass the channel as an argument
	go omnicrementing(myChan)   // call the function in a goroutine

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
