package main

import "fmt"

func main() {
	// -----> POINTERS <-----

	// slices and maps works with pointers and could change when psevdo copying

	var a  int = 40
	var b *int = &a
	var c *int = b
	// b holds the memory location that is holding the data, so the address
	// if we want to dereference using *b
	fmt.Println(a, b)
	a = 20
	fmt.Println(a, *b)
	fmt.Println(a, *b, *c)

	// we can take the address with new()
	// ugly
	var ms *animal
	ms = new(animal)
	ms.name = "ivan" // not allowed in other language like c++, it is wrong
	// ms.name is clean of (*ms).name, go helps us
	fmt.Println(ms)
}

type animal struct {
	name string
}


