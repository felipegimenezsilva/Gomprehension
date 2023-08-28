package slice

import (
	"Gomprehension/slice"
	"testing"
)

func initIntSlice() []int {
	var mySlice []int
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 8)
	return mySlice
}

func testMapFunction(x int) int {
	return x * 4
}

func testFilterFunction(x int) bool {
	return x%2 == 0
}

func TestMap(t *testing.T) {
	var test = initIntSlice()
	var result = slice.Map(test, testMapFunction)
	for index := 0; index < len(test); index++ {
		if result[index] != testMapFunction(test[index]) {
			t.Error("slice.Map: the expected value does not match with the map")
		}
	}
}

func existInSlice[T comparable](slice []T, value T) bool {
	for index := 0; index < len(slice); index++ {
		if slice[index] == value {
			return true
		}
	}
	return false
}

func TestFilter(t *testing.T) {
	var test []int = initIntSlice()
	var result = slice.Filter(test, testFilterFunction)
	var count int = 0
	for index := 0; index < len(test); index++ {
		var keepElement bool = testFilterFunction(test[index])
		if keepElement {
			count++
		}
		if keepElement != existInSlice[int](result, test[index]) {
			t.Error("slice.Filter: the filtered list is incorrect")
		}
	}
	if count != len(result) {
		t.Error("slice.Filter: length of filtered list is incorrect")
	}
}

func TestReduce(t *testing.T) {
	var test []int = initIntSlice()
	var result int = slice.Reduce[int](test, func(i1, i2 int) int {
		return i1 + i2
	})
	var sum int = 0
	for _, value := range test {
		sum += value
	}
	if sum != result {
		t.Error("slice.Reduce: unexpected result obtained")
	}

}
func TestMax(t *testing.T) {
	var test []int = initIntSlice()
	var max int = test[0]
	for _, item := range test {
		if max < item {
			max = item
		}
	}
	var maxIndex int = slice.Max[int, int](test, func(i1 int) int {
		return i1
	})
	var result = test[maxIndex]
	if result != max {
		t.Error("slice.Max: unexpected max value")
	}
}
func TestMin(t *testing.T) {
	var test []int = initIntSlice()
	var min int = test[0]
	for _, item := range test {
		if min > item {
			min = item
		}
	}
	var minIndex int = slice.Min[int, int](test, func(i1 int) int {
		return i1
	})
	var result = test[minIndex]
	if result != min {
		t.Error("slice.Min: unexpected min value")
	}
}

func TestZip(t *testing.T) {}

// Felipe Gimenez
