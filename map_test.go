package itergo_test

import (
	"github.com/azvaliev/itergo"
	"github.com/azvaliev/itergo/testutils"
	"testing"
)

func TestKeepIdentical(t *testing.T) {
	expectedSlice := []int{5, 6, 7, 8}
	gotSlice := itergo.Map(expectedSlice, func(item int, idx int) int {
		return item
	})

	if err := testutils.SliceDeepEqual(expectedSlice, gotSlice); err != nil {
		t.Error(err)
	}
	if &expectedSlice == &gotSlice {
		t.Error("The original slice and the copy are the same")
	}
}

func TestMapToIndices(t *testing.T) {
	testSlice := []string{"a", "b", "c", "d"}

	expectedSlice := []int{0, 1, 2, 3}
	gotSlice := itergo.Map(testSlice, func(item string, idx int) int {
		return idx
	})

	if err := testutils.SliceDeepEqual(expectedSlice, gotSlice); err != nil {
		t.Error(err)
	}
}

func TestMapStructProperty(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	testSlice := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	expectedSlice := []string{"Alice", "Bob", "Charlie"}
	gotSlice := itergo.Map(testSlice, func(item Person, idx int) string {
		return item.Name
	})

	if err := testutils.SliceDeepEqual(expectedSlice, gotSlice); err != nil {
		t.Error(err)
	}
}

func TestMapArray(t *testing.T) {
	testSlice := [3]int{1, 2, 3}

	expectedSlice := [3]int{2, 4, 6}
	gotSlice := itergo.Map(testSlice[:], func(item int, idx int) int {
		return item * 2
	})

	if err := testutils.SliceDeepEqual(expectedSlice[:], gotSlice); err != nil {
		t.Error(err)
	}
}
