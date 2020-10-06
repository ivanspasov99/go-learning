package main

import "fmt"

// varargs
func sum(values ...int) int {
	var sum = 0
	for _, v := range values {
		sum += v
	}
	return sum
}

// we can return pointer, the go recognise that we are returning values so it promote the variable to be on shared memory
func sumPoint(values ...int) *int {
	var sum = 0
	for _, v := range values {
		sum += v
	}

	var pointer *int = &sum
	return pointer

	// or
	// return &sum
}

// declaring that we are going to return result int, so result variable
// is automatically returned and available in the body of the function
// confusing
func strangeReturnValue(len, cap int) (result int) {
	result = len + cap
	return
}

// multiple return values
// used with errors
func multipleReturnPar(a, b int) (int, error) {
	// check for corner case and return error object
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	// if everything is ok...
	return a / b, nil
}

// methods
type human struct {
	name string
}

// funct (human human) getName() {
// code...
// }

// bu that not change the real value, it is making copy, so we need to pass reference
func (human human) changeName(name string) {
	human.name = name
}

func (human *human) changeOriginalName(name string) {
	human.name = name
}

func main() {

	// anonymous function
	a := 10
	func(par int) {
		fmt.Println("Execute immediately")
	}(a) // here is the parameter


	fmt.Println(sum(1, 2, 3))
}
