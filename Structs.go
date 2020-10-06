package main

import (
	"fmt"
	"reflect"
)

// -----> STRUCT <-----
// unlike maps and slice structs are referring to different data, so when copy, there is newly create struct
// that can be memory ineffective


// example
type Doctor struct {
	number int
	name string
	companions []string
	cat Pet

}

type Pet struct {
	name string
	gender string
}

// instead of INHERITANCE
// COMPOSITION RELATIONSHIP
// doing with embedded structs
type Animal struct {
	name string
}
type Bird struct {
	Animal // this gets us all Animal fields, so in this case 'name' only
	kind string
}


// concept tags -> specific requirements
// reflect package
// used for validation for example
// if there is this kind of tag, so do that....
type Fish struct {
	name string `length:"100"`
	origin string
}

var Cat = Pet {
name: "sara",
gender: "man",
}

func main() {

	t := reflect.TypeOf(Fish{})
	field, _ := t.FieldByName("name")
	fmt.Println(field)

	doc := Doctor{
		number: 1,
		name: "Ivan",
		companions: []string {
			"Petar",
			"Gosho",
		},
		cat: Cat,
	}

	var temp = Doctor{1, "name", []string{"asd"}, Pet{"cat", "cat"}}
	fmt.Println(doc, temp)

	// rarely used
	man := struct{name string} {name: "Name"}
	fmt.Println(man)
}
