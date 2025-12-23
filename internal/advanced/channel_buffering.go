package advanced

import "fmt"

func IntroChannelBuffered() {
	//* Create channel
	ch := make(chan int)
	//* Send a value to the channel
	//* We need goroutines for unbufferd channels and not for bufferd channels
	//* Unbuffered channels immediately look for a receiver
	//* We need the sender and the receiver at the same time
	go func() {
		ch <- 1
	}()
	//* How is the go runtime working with channels and goroutines
	//? 1) Initiallize a channel
	//? 2) Start a go routine (extracts the func out of the main thread moves to next line)
	//? at the same time another thread excecutes the ( ch <- 1 ) and it moves to that line
	//? and it tells it that we are receiving the value of one
	//? 3) The Go runtime checks that there is a receiver and it transfers the value to the variabl
	//* Receive value from channel into receiver
	receiver := <-ch
	fmt.Println(receiver)

	//*** Unbuffered channels block on receive if there is no corresponding send operation ready and as soon as there is a send operation then it doesnt block
	//* Unbuffered channels also block on send if there is no corresponding receive operation

	//! Changing subject to Buffered Channels
	//* Buffered channels provide ->
	//* 1) Async Communication
	//* 2) Load Balancing
	//* 3) Flow Control

	//* Creating Buffered Channel
	//* 1) make(ch int, 8) // make(channel type, capacity)
	//* 2) Buffer Capacity
	// bufCh := make(chan int, 10)

	//* Key concepts of Channel Buffering
	//* 1) Blocking Behavior
	//* 2) Non-Blocking Operations
	//* 3) Impact on Performance

	//* Best Practice for Using Buffered Channels
	//* 1) Avoid Over-Buffering
	//* 2) Graceful Shutdown
	//* 3) Monitoring Buffer Usage
}
