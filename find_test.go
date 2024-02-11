package itergo

import "testing"

func TestFindEmpty(t *testing.T) {
	var empty []int
	_, idx := Find(empty, func(int, int) bool { return true })
	if idx != -1 {
		t.Errorf("Expected -1, got %d", idx)
	}

	_, idx = FindLast(empty, func(int, int) bool { return true })
	if idx != -1 {
		t.Errorf("Expected -1, got %d", idx)
	}
}

func TestFind(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}

	item, idx := Find(collection, func(item int, idx int) bool { return item%2 == 0 })
	if item != 2 || idx != 1 {
		t.Errorf(
			"Expected item %d at index %d, got item %d at index %d",
			2,
			1,
			item,
			idx,
		)
	}

	item, idx = FindLast(collection, func(item int, idx int) bool { return item%2 == 0 })
	if item != 4 || idx != 3 {
		t.Errorf(
			"Expected item %d at index %d, got item %d at index %d",
			4,
			3,
			item,
			idx,
		)
	}
}
