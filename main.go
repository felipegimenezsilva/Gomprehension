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

	fmt.Println(array)

	fmt.Println(mapVals)

	var filtered []int = slice.Filter(array, func(item int) bool {
		return (item % 2) == 0
	})

	fmt.Println(filtered)

	//var mm = slice.MinMax(array, func(item int) int {
	//	return item * 2
	//})
	//
	//fmt.Println(mm)

	var red int = slice.Reduce(array, func(a int, b int) int {
		return a + b
	})

	fmt.Println(red)

	var temp []int
	temp = append(temp, 3)
	temp = append(temp, 2)
	temp = append(temp, 1)

	var x [][]int = slice.Zip(array, array, array, temp)
	fmt.Println(x)
}
