package main

import (
	"fmt"
	"time"
)

func main() {

	// simple channel

	incrementChannel := make(chan int)

	// run goroutine
	go SendMessage(incrementChannel)

	// read from channel
	for i := 0; i < 5; i++ {
		msg := <-incrementChannel
		fmt.Println("increment channel:", msg)
	}
	/*
		for msg := range incrementChannel {
			fmt.Println("increment channel:", msg)
		}
	*/

	// asynchronous buffered channel

	asyncChannel := make(chan int, 5)

	// write to channel
	for i := 0; i < 5; i++ {
		asyncChannel <- i // can be locked here, waiting for readers
	}

	// close writting, you cannot write after close
	close(asyncChannel)

	// read from channel
	for i := 0; i < 5; i++ {
		fmt.Println("async buffered channel:", <-asyncChannel)
	}

	// select between channels (get data from several channels)

	// create simple ticker with its channel
	ticker := time.NewTicker(time.Second)

	for i := 0; i < 5; i++ {
		select {
		case incMsg := <-incrementChannel:
			fmt.Println("select increment channel:", incMsg)
		case tickMsg := <-ticker.C:
			fmt.Println("select ticker channel:", tickMsg)
		}
	}
}

// SendMessage sends message to channel
func SendMessage(c chan<- int) {
	increment := GenerateCounter()

	for { // while true, infinite writing
		time.Sleep(time.Second)
		c <- increment()
	}
}

// GenerateCounter generates incremental counter
func GenerateCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
