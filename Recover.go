package main

import "fmt"

// -----> RECOVER <-----

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {

	// Recover is a built-in function that regains control of a panicking goroutine.
	// Recover is only useful inside deferred functions.
	// useful in defer because it is executed before the panic
	// in this scenario it will return up to the stack to the function f after panic in g
	// and will execute defer, in this defer we recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	panic("This is the value that is stored in 'r'")
	// g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
