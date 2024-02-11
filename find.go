package itergo

type FindCondition[T any] func(item T, idx int) bool

// Find returns the first element in a slice that meets the condition function.
// If no elements meet the condition or the slice is empty, the zero value for the
// element will be returned, and -1 for the index
func Find[T any](collection []T, condition FindCondition[T]) (item T, idx int) {
	for idx, item := range collection {
		if condition(item, idx) {
			return item, idx
		}
	}

	var notFound T
	return notFound, -1
}

// FindLast is similar to Find, except it returns the last element and it's index
// which meets the condition
func FindLast[T any](collection []T, condition FindCondition[T]) (item T, idx int) {
	var found T
	foundIdx := -1

	for idx, item := range collection {
		if condition(item, idx) {
			found = item
			foundIdx = idx
		}
	}

	return found, foundIdx
}
