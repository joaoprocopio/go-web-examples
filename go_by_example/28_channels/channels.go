package main

import (
	"fmt"
	"time"
)

// channels are the pipes that connect concurrent goroutines
// you can send values into channels from one goroutine and receive those values into another goroutine

func main() {
	// create a new channel with make(chan val-type)
	// channels are typed by the values they convey
	messages := make(chan string)

	// send a new value int oa channel using the channel <- syntax
	// here we send "ping" to the messages channel we made above
	// from a new goroutine
	go func() {
		time.Sleep(time.Second * 5)
		messages <- "ping"
	}()

	// the <-channel syntax receives a value from the channel
	// here we'll receive the "ping" message we sent above and print it out
	// note: 	this syntax makes the current execution blocks until the channel receives a value
	// 			if the coroutine that you're channeling to doesn't receives a value
	// 			the current execution deadlocks
	// 			https://www.craig-wood.com/nick/articles/deadlocks-in-go/
	msg := <-messages
	fmt.Println(msg)
}

// when we run the program the "ping" message is successfully passed from one goroutine to another via our channel

// by default sends and receives block until both the sender and receiver are ready
// this property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization
