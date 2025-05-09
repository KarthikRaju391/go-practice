// ---- This is an implementation of the done channel -----
package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing work!")
		}
	}
}

func main() {
	// creating a done channel
	done := make(chan bool)

	// pass the done channel to the goroutine
	go doWork(done)

	// let the indefinite goroutine work for 5 seconds
	time.Sleep(time.Second * 5)

	// then, close the channel
	close(done)

	fmt.Println("Hello, world!")
}
