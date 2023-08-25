package main

import (
	"Gomprehension/slice"
	"fmt"
)

func testMap(item int) int {
	return item * item
}

func main() {

	var array []int
	array = append(array, 1)
	array = append(array, 2)
	array = append(array, 3)
	array = append(array, 4)
	array = append(array, 5)

	var mapVals = slice.Map(array, testMap)

	fmt.Print(mapVals)

}
