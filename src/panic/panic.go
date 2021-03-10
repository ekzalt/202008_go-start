package main

import (
	"errors"
	"fmt"
	"runtime"
)

func main() {
	// try-catch in go - panic-recover
	ThrowPanic()
}

// ThrowPanic throws panic
func ThrowPanic() {
	// register catcher of panic inside a dangerous function
	defer CatchPanic()

	panic("Something went wrong")
}

// CatchPanic catches panic
func CatchPanic() {
	// check if panic occurred
	if recovered := recover(); recovered != nil {
		fmt.Println("panic detected")

		// capture stack trace
		buffer := make([]byte, 10000)
		runtime.Stack(buffer, false)
		fmt.Println("panic stack trace", string(buffer))
	}
}

// SomeFunc is a function with error handling inside
func SomeFunc() {
	// this is catching of error
	defer func() {
		if recover() != nil {
			err := errors.New("Some error text here")
			fmt.Println(err)
		}
	}()

	// some important logic here, can thow error
}
