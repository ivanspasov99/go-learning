package main

import "fmt"

func main() {
	  // -----> DEFER <-----
	 // defer is executing after a function is done (when returned)
	 // defer key usage - when closing channels, files etc
	 // res.open()
	 // defer res.close()
	 // a lot ... code

	// USAGE - defer takes arguments values compile time
	// if we change the value after that it won't change nothing in defer
	a := 10
	defer fmt.Println(a) // print 10, cause compile time it is 10
	// the value 5 is evaluated compile time and in 15 row it is 10, so evaluated to 10
	a = 5

	// -----> PANIC <-----
	// PANIC happens after defer statements
	// when program see that defer exist, executes is and then panic
	// used because something could be staying open like file...

	// We’ll use panic to check for unexpected errors
	// panic("There is a problem")
	//a, b := 1, 0
	//ans := a / b
	//fmt.Println(ans)

}
