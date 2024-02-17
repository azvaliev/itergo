package itergo

// Contains checks if a slice contains a specific item, returns true if it does, false otherwise.
func Contains[T comparable](collection []T, wantItem T) bool {
	for _, item := range collection {
		if item == wantItem {
			return true
		}
	}

	return false
}
