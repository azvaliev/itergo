package itergo_test

import (
	"github.com/azvaliev/itergo"
	"math/big"
	"testing"
)

func TestAllCorrectOrderAndIndex(t *testing.T) {
	expectedSlice := []int{1, 2, 3, 4}
	_ = itergo.All(expectedSlice, func(item int, idx int) bool {
		if item != expectedSlice[idx] {
			t.Errorf("Expected %d at position %d but got %d", expectedSlice[idx], idx, item)
		}
		return true
	})
}

func TestAllEmptySlice(t *testing.T) {
	_ = itergo.All([]int{}, func(item int, idx int) bool {
		t.Error("Slice is empty, should not reach here")
		return false
	})
}

func TestAllOneFalse(t *testing.T) {
	expectedSlice := []int{1, 2, 3, 4}
	if itergo.All(expectedSlice, func(item int, idx int) bool {
		return item != 3
	}) {
		t.Error("One of the items is not 3")
	}
}

func TestAllTrue(t *testing.T) {
	expectedSlice := []string{"a", "b", "c", "d"}

	if !itergo.All(expectedSlice, func(item string, idx int) bool {
		return true
	}) {
		t.Error("Should not reach here, all true")
	}
}

func TestAllFalse(t *testing.T) {
	expectedSlice := []string{"a", "b", "c", "d"}

	if itergo.All(expectedSlice, func(item string, idx int) bool {
		return false
	}) {
		t.Error("Should not reach here, all false")
	}
}

func TestAllPrimeNumbers(t *testing.T) {
	expectedSlice := []int{2, 3, 5, 7, 11, 13, 17, 19}

	if !itergo.All(expectedSlice, func(item int, idx int) bool {
		return big.NewInt(int64(item)).ProbablyPrime(0)
	}) {
		t.Error("Should not reach here, all prime numbers")
	}
}

func TestAllArray(t *testing.T) {
	expectedArray := [4]string{"a", "b", "c", "d"}

	if !itergo.All(expectedArray[:], func(item string, idx int) bool {
		return true
	}) {
		t.Error("Should not reach here, all true")
	}
}
