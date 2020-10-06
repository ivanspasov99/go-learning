package main

import "fmt"

func main() {
	// -----> MAPS <-----
	// passed by reference

	pop := map[string]int {
		"Ivan": 18,
		"Magi": 15,
		"Hagi": 17,
	}

	fmt.Println(pop)
	// add
	name := "Denis"
	pop[name] = 20
	// delete
	delete(pop, "Hagi")

	//for s := range pop {
	//	fmt.Println("Key:", s, "Value:", pop[s])
	//}

	for key, value := range pop {
		fmt.Println(key, value)
	}

	// comma, ok syntax
	// map return two values: value and true/false if the returned value exist
	key, ok := pop["notFound"]
	fmt.Println(key, ok)

	// _ -> use it because we want only the second return value
	_, ok = pop["Magi"]
	fmt.Println(ok)

	//fmt.Println("Len: ", len(pop))
	//fmt.Println(pop)
}
