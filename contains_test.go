package itergo_test

import (
	"github.com/azvaliev/itergo"
	"testing"
)

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
