package main

import (
	"fmt"
)

func main() {
	// -----> ARRAYS <-----
	// use slice -> better
	grades := [3]int{6, 5, 4}

	for i:=0; i < len(grades); i++ {
		fmt.Printf("value = %v\n", grades[i])
	}

	// ... -> means that the size of arrays is as much as the elements are
	marks := [...]int{1, 2, 3}
	fmt.Println(marks)


	var arr [3]string
	arr[0] = "ivan"
	fmt.Println(arr)

	var twoDism = [][]int {{1, 1}, {2,2}}
	//printTwoDism(2, 2, twoDism)


	// when copying arrays, you are copying array not just where it points to (like in C++ pointers)
	copyArr := twoDism
	twoDism[0][0] = 100
	printTwoDism(2, 2, copyArr)


}
// !!! the copying array is slow when it is big, so use slice....
func printTwoDism(row int, col int, arr [][]int) {
	for i:=0; i<row; i++ {
		fmt.Println()
		for j:=0; j<col; j++{
			fmt.Printf("%v ", arr[i][j])
		}
	}
}
