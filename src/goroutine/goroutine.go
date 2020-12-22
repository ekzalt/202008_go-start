package main

import (
	"fmt"
	"sync"
)

func main() {
	// mutex

	user := User{new(sync.RWMutex), "Max"}
	user.SetName("Den")
	fmt.Println("User:", user.GetName())

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
	*sync.RWMutex
	Name string
}

// SetName sets user name with mutex
func (u *User) SetName(name string) {
	u.Lock()
	defer u.Unlock()

	// modify locked object
	u.Name = name

	// u.Unlock()
}

// GetName gets user name with mutex
func (u *User) GetName() string {
	// read lock prevents writing during reading
	u.RLock()
	defer u.RUnlock()

	return u.Name
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
