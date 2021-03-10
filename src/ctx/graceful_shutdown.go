package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func handleSignals(callback context.CancelFunc) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	for {
		sig := <-sigChan

		switch sig {
		case os.Interrupt:
			callback()
			return // break loop and return from function
		}
	}
}

func startServer(ctx context.Context) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")

	if err != nil {
		return err
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		return err
	}

	defer tcpListener.Close()

	for {
		// listen for done event from ctx done channel
		select {
		case <-ctx.Done(): // read-only channel
			fmt.Println(ctx.Err()) // context canceled
			return nil
		default:
			// set timeout of waiting for a new connection 1 second
			if err := tcpListener.SetDeadline(time.Now().Add(time.Second)); err != nil {
				return err
			}

			// wait for connection
			conn, err := tcpListener.Accept()

			if err != nil {
				if os.IsTimeout(err) {
					continue
				}

				return err
			}

			fmt.Println("new connection received", conn.RemoteAddr().String())
		}
	}
}

func main() {
	ctx, finish := context.WithCancel(context.Background())
	go handleSignals(finish)

	if err := startServer(ctx); err != nil {
		log.Fatal(err)
	}
}
