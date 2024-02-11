package itergo

// Filter applies a condition function to each item in a slice and returns a
// new slice with the items that pass the condition. (When the condition function returns true)
func Filter[T any](collection []T, condition func(item T, idx int) bool) []T {
	filteredCollection := make([]T, 0)

	for idx, item := range collection {
		if condition(item, idx) {
			filteredCollection = append(filteredCollection, item)
		}
	}

	return filteredCollection
}
