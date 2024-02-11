package itergo

import (
	"github.com/azvaliev/itergo/testutils"
	"testing"
)

func TestFilterAllTrue(t *testing.T) {
	expectedSlice := []int{1, 2, 3, 4, 5}

	gotSlice := Filter(expectedSlice, func(item int, idx int) bool {
		return item > 0
	})

	if err := testutils.SliceDeepEqual(expectedSlice, gotSlice); err != nil {
		t.Error(err)
	}
}

func TestFilterAllFalse(t *testing.T) {
	expectedSlice := []int{1, 2, 3, 4, 5}

	gotSlice := Filter(expectedSlice, func(item int, idx int) bool {
		return false
	})

	if err := testutils.SliceDeepEqual([]int{}, gotSlice); err != nil {
		t.Error(err)
	}
}

func TestFilterEmptySlice(t *testing.T) {
	gotSlice := Filter([]int{}, func(item int, idx int) bool {
		return true
	})

	if err := testutils.SliceDeepEqual([]int{}, gotSlice); err != nil {
		t.Error(err)
	}
}

func TestFilterStructProperty(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	testSlice := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	expectedSlice := []Person{
		{"Alice", 25},
		{"Bob", 30},
	}
	gotSlice := Filter(testSlice, func(item Person, idx int) bool {
		return item.Age < 35
	})

	if err := testutils.SliceDeepEqual(expectedSlice, gotSlice); err != nil {
		t.Error(err)
	}
}
