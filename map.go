package itergo

// Map applies a mapper/transformer function to each item in a slice
// and returns a new slice with the mapped items.
func Map[T any, U any](collection []T, mapper func(item T, idx int) U) []U {
	mappedCollection := make([]U, len(collection))
	for idx, item := range collection {
		mappedCollection[idx] = mapper(item, idx)
	}

	return mappedCollection
}
