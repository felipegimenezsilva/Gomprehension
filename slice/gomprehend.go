package slice

type Numerical interface {
	~int | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// apply a function for each element in slice
func Map[T1, T2 any](listItems []T1, function func(T1) T2) []T2 {
	var result []T2
	for _, item := range listItems {
		result = append(result, function(item))
	}
	return result
}

// remove items that does not match with some criteria
func Filter[T any](listItems []T, function func(T) bool) []T {
	var result []T
	for _, item := range listItems {
		if function(item) {
			result = append(result, item)
		}
	}
	return result
}

// apply some function and accumule the values
func Reduce[T any](listItems []T, function func(T, T) T) T {
	if len(listItems) < 2 {
		panic("Reduce: itemList must have at least 2 elements")
	}
	var result T = listItems[0]
	for index := 1; index < len(listItems); index++ {
		result = function(result, listItems[index])
	}
	return result
}

// search for the maximum value in a list, and return the index
func Max[T any, N Numerical](itemList []T, function func(T) N) int {
	var maxIndex int = 0
	var maxValue N = function(itemList[0])
	for index := 1; index < len(itemList); index++ {
		var currentValue N = function(itemList[index])
		if currentValue > maxValue {
			maxIndex = index
			maxValue = currentValue
		}
	}
	return maxIndex
}

// search for the minimum value in a list, and return the index
func Min[T any, N Numerical](itemList []T, function func(T) N) int {
	var minIndex int = 0
	var minValue N = function(itemList[0])
	for index := 1; index < len(itemList); index++ {
		var currentValue N = function(itemList[index])
		if currentValue < minValue {
			minIndex = index
			minValue = currentValue
		}
	}
	return minIndex
}

// similar to python zip function
func Zip[T any](listItems ...[]T) [][]T {
	var listSizes []int = Map(listItems, func(items []T) int {
		return len(items)
	})
	var size int = Min(listSizes, func(item int) int { return item })
	var result [][]T
	for index := 0; index < size; index++ {
		var group []T
		for itemIndex := 0; itemIndex < len(listItems); itemIndex++ {
			group = append(group, listItems[itemIndex][index])
		}
		result = append(result, group)
	}
	return result
}

// Felipe Gimenez
