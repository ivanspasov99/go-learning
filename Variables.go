package main

import (
	"fmt"
	"strconv"
)

// we can declare package variables
var packageVar = 10

// variable visibility
var v = 5 // visible inside the package, starts with lowercase letter (UNEXPORTED)
var V = 10 // visible outside the package (EXPORTED)

// var blocks - clean view
var (
	name = "ivan"
	lastName = "spasov"
)

func main() {
	// -----> Declaring Variables <-----

	// uninitialized variable always has:
		// zero if number
		// empty string if string
		// false if boolean

	var zero int
	var str1 string
	var bo bool
	fmt.Println(zero)
	fmt.Println(str1)
	fmt.Println(bo)

	// FIRST
	var i int
	i = 42
	fmt.Println("i =", i)

	// SECOND
	// int is redundant because we assign integer value to it
	var b int = 105
	fmt.Println("b =", b)

	// THIRD
	// VALID ONLY IN FUNCTIONS, NOT GLOBALLY
	c := "string"
	fmt.Println("c =", c)

	// Package Variable
	fmt.Println("package =", packageVar)


	// -----> CASTING <-----

	var k = float32(i)
	fmt.Println("k float32 =", k)

	// convert to string -> needs

	str := strconv.Itoa(b)
	fmt.Printf("Value: %v  Type: %T\n", str,  str)

}
