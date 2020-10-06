package main

import "fmt"

func main() {
	// -----> SLICE <-----
	// passed by reference

	// operation on one slice is affecting all other that are pointing to him
	a := []int{1, 2, 3, 4}
	fmt.Printf("Len: %v ", len(a))
	fmt.Printf("Cap: %v\n", cap(a))

	// all of them reference to original array
	// new slices
	//b := a[:]
	//c := a[1:]
	//d := a[:2]
	//e := a[3:4]

	currentElements := 0
	capacity := 0
	b := make([]int, currentElements, capacity)
	fmt.Println(b)
	// dynamically  increase capacity of the b slice
	b = append(b, 1, 2, 3, 4)
	fmt.Printf("Len: %v ", len(b))
	fmt.Printf("Cap: %v\n", cap(b))

	// remove elements form top and bottom
	c := b[1:]
	c[1] = 10
	fmt.Println(c)
	fmt.Println(b)

	// remove elements form mid
	g := append(b[:2], b[3:]...)
	fmt.Println(g)
	fmt.Println(b)



}
