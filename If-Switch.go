package main

import "fmt"

func main() {
	pop := map[string]int {
		"ivan": 1,
		"hah": 2,
	}

	// initializer -> effective go, variables only visible in the scope
	if value, ok := pop["ivan"]; ok {
		fmt.Println(value)
	}

	// switch
	// fallthrough -> keyword to continue to the next statement, executes it no matter the if/switch clause
	var t interface{} = 5 // interface type is the parent of everything, everything is interface
	// so we can do stuff like that, to check what type is coming
	// check if type milk, or potato etc.

	switch x := t.(type) {
	case int: {
		fmt.Println("i am int", x)

	}
	case string: {
		fmt.Println("i am string", x)
	}
	case float64: {
		fmt.Println("i am struct", x)
	}
	}
}
