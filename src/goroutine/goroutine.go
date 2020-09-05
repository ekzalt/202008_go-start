package main

import (
	"fmt"
	"sync"
)

func main() {
	// mutex

	user := User{new(sync.Mutex), "Max"}
	SetUserName(&user, "Den")

	// goroutine

	wg := new(sync.WaitGroup)
	// make 5 threads
	for i := 0; i < 5; i++ {
		go Process(i, wg)
	}
	// wait for all call wg.Done()
	wg.Wait()

	once := sync.Once{}
	// the anonymous function will be called only once
	for i := 0; i < 5; i++ {
		once.Do(func() { fmt.Println("this should call only once") })
	}
}

// mutex

// User sync with Mutex
type User struct {
	NameMutex *sync.Mutex
	Name      string
}

// SetUserName set user name with mutex
func SetUserName(u *User, name string) {
	u.NameMutex.Lock()
	defer u.NameMutex.Unlock()

	// modify locked object
	u.Name = name

	// u.NameMutex.Unlock()

	fmt.Println("User:", *u)
}

// goroutine

// Process prints i
func Process(i int, wg *sync.WaitGroup) {
	wg.Add(1) // number of sub-threads
	defer wg.Done()

	// do some hard work
	fmt.Println("run goroutine", i)

	// wg.Done()
}
