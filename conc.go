package main

import "fmt"

func main() {
	// createing a buffered channel that holds 10 funcs
	cnp := make(chan func(), 10)

	// worker pool
	// 4 workers wait for the work here
	for i := 0; i < 4; i++ {
		go func() {
			// range over the channel for recieving the tasks
			for f := range cnp {
				// if worker gets the task execute it
				f()
			}
		}()
	}

	// we send the task to the buffered channel
	cnp <- func() {
		// Heremain exits before workers can process
		fmt.Println("HERE1")
	}
	// main function continues ..
	fmt.Println("Hello")

}
