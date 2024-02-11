package itergo

// All returns true if all elements in the slice satisfy the condition,
// if any element does not satisfy the condition, it returns false.
func All[T any](collection []T, condition func(item T, idx int) bool) bool {
	for idx, item := range collection {
		if !condition(item, idx) {
			return false
		}
	}

	return true
}
