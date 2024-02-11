package testutils

import "fmt"

// SliceDeepEqual verify two slices have the same item values in the same order
func SliceDeepEqual[T comparable](wantSlice []T, gotSlice []T) error {
	if len(wantSlice) != len(gotSlice) {
		return fmt.Errorf("want length of %d, got %d", len(wantSlice), len(gotSlice))
	}

	for i := 0; i < len(wantSlice); i++ {
		if wantSlice[i] != gotSlice[i] {
			return fmt.Errorf(
				"want %#v at position %d, got %#v",
				wantSlice[i],
				i,
				gotSlice[i],
			)
		}
	}

	return nil
}
