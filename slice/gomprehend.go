package slice

func Map[T any](listItems []T, function func(T) T) []T {
	var result []T
	for _, item := range listItems {
		result = append(result, function(item))
	}
	return result
}

func Filter[T any](listItems []T, function func(T) bool) []T {
	var result []T
	for _, item := range listItems {
		if function(item) {
			result = append(result, item)
		}
	}
	return result
}
