package main

import "fmt"


// iota examples -> can be used as enums: check -> https://yourbasic.org/golang/iota/
const (
	a = iota
	_ // skip value, do not get the value, used when we do not want the value
	b = iota
	c = iota
)
// alternative is
//const (
//	a = iota
//	b
//	c
//)

func main() {
	// -----> PRIMITIVES <-----

	var t = true
	f := false

	fmt.Println(f)
	fmt.Println(t)

	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v %T\n", n, n)
	fmt.Printf("%v %T", m, m)

	// -----> CONSTANTS <-----

	// not uppercase because of exported vars
	const myConst = 42
	fmt.Println(myConst)

	// it is not working - incompatible types
	//	var a = 42
	//	var b int16 = 20
	//	fmt.Println(a + b)

	// ... but with constants it is working
	// constants are directly replaced with numbers on compile time so they can be seen as the needed type (in this case int16)
	const q = 42
	var w int16 = 20
 	fmt.Println(q + w)

	// iota -> increments constants values that are in const (...values) until another constant (then it resets)
	fmt.Println(a, b, c)
	const stopIota = 10
	const d = iota
	fmt.Println(a, b, c, d)

}

