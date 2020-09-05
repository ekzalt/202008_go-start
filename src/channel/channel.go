package main

import (
	"fmt"
	"time"
)

func main() {

	// simple channel

	incrementChannel := make(chan int) // buffer 1 by default
	defer close(incrementChannel)

	// run goroutine
	go SendMessage(incrementChannel)

	// read from channel

	// option 1 - iterate with simple for
	for i := 0; i < 5; i++ {
		msg := <-incrementChannel
		fmt.Println("increment channel:", msg)
	}

	/*
		// option 2 - iterate with range
		for msg := range incrementChannel {
			fmt.Println("increment channel:", msg)
		}
	*/

	// asynchronous buffered channel

	// async channel with buffer 3
	asyncChannel := make(chan int, 3)
	// async write with goroutine to channel, 6 elems to 3 buffered channel
	go Produce(asyncChannel)
	// read from channel
	Consume(asyncChannel)

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

// Produce writes 6 ints to 3 buffered async channel
func Produce(ch chan int) {
	for i := 0; i < 6; i++ {
		ch <- i // can be locked here, waiting for readers
	}

	// close writting, you cannot write after close
	close(ch)
}

// Consume reads all ints from buffered channel
func Consume(ch chan int) {
	/*
		// option 1 - iterate with simple for
		for i := 0; i < 6; i++ {
			fmt.Println("async buffered channel:", <-asyncChannel)
		}
	*/

	// option 2 - iterate with range
	for n := range ch {
		fmt.Println("async buffered channel:", n)
	}
}
