package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ReadFromStdin()
}

// ReadFromStdin reads and prints name from stdin
func ReadFromStdin() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("How is your name?")
		text, _ := reader.ReadString('\n')
		fmt.Println("Hello", text)
	}
}
