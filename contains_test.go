package itergo_test

import (
	"fmt"
	"github.com/azvaliev/itergo"
	"testing"
)

func ExampleContains() {
	// Let's check if this slice contains the number 3
	slice := []int{1, 2, 3, 4, 5}

	containsThree := itergo.Contains(slice, 3)
	if containsThree {
		fmt.Println("The slice contains the number 3")
	} else {
		fmt.Println("The slice does not contain the number 3")
	}
	// Output: The slice contains the number 3
}

func TestContains(t *testing.T) {
	shouldHaveThree := itergo.Contains([]int{1, 2, 3, 4, 5}, 3)
	if !shouldHaveThree {
		t.Error("Should have 3")
	}
}

func TestContainsArray(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	shouldHaveThree := itergo.Contains(arr[:], 3)

	if !shouldHaveThree {
		t.Error("Should have 3")
	}
}

func TestContainsEmpty(t *testing.T) {
	shouldHaveThree := itergo.Contains([]int{}, 3)
	if shouldHaveThree {
		t.Error("Should not have 3")
	}
}

func TestNotContains(t *testing.T) {
	shouldHaveThree := itergo.Contains([]int{1, 2, 4, 5}, 3)
	if shouldHaveThree {
		t.Error("Should not have 3")
	}
}

func TestContainsMultiple(t *testing.T) {
	shouldHaveThree := itergo.Contains([]int{1, 2, 3, 4, 5, 3}, 3)
	if !shouldHaveThree {
		t.Error("Should have 3")
	}
}
