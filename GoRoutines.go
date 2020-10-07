package main

import (
	"fmt"
	"strconv"
	"sync"
)

// -----> GOROUTINES <-----

var s = sync.WaitGroup{}
var m = sync.RWMutex{}

var counter = 0
func main() {

	// BASIC EXAMPLE
	//basic()

	// SYNCHRONISATION
	//var wg = sync.WaitGroup{}
	//wg.Add(1)
	//go func () {
	//	sayHello("Sync")
	//	wg.Done()
	//}()
	//wg.Wait()

	// Multiple synchronisation (WaitGroup) but here there is race between sayHello and sayGoodbye, we do not which is going to be first
	// so introducing RWMutex{} - read & write mutex
	for i := 0; i < 10; i++ {
		//s.Add(2)
		m.RLock() // RLock(): multiple go routine can read(not write) at a time by acquiring the lock.
		go read("Iteration " + strconv.Itoa(i))
		m.Lock() // Lock(): only one go routine read/write at a time by acquiring the lock.
		go write("Iteration " + strconv.Itoa(i))
		//s.Wait()
	}
}

func read(data string)  {
	fmt.Println("Data:", counter)
	//s.Done()
	m.RUnlock()
}

func write(data string) {
	// we write something
	counter++
	m.Unlock()
	//s.Done()
}


func basic() {
	// go routines, run function in greenThreads
	// spawn new go routine, and starting new process which is different from main()
	// so we need to wait for go sayHello() to be executed before proceeding
	go read("Without Sync")

	// it can access var declared in upper scopes
	var msg = "Anonymous functions"
	go func() {
		fmt.Println(msg)
	}()
	// creates race condition, you do not know which will access msg var first
	msg = "Main is faster than anonymous function"
}