package itergo_test

import (
	"fmt"
	"github.com/azvaliev/itergo"
	"testing"
)

func ExampleFind() {
	// Let's say we have an array of nodes and we want to find a node with a certain ID
	type Node struct {
		ID       int
		Children []int
	}
	array := []Node{
		{1, []int{2, 3}},
		{2, []int{4, 5}},
		{3, []int{6, 7}},
		// ...
	}

	// In the filter function, we return true if the node's ID is 3
	node, idx := itergo.Find(array, func(node Node, idx int) bool {
		return node.ID == 3
	})

	// If the node is found, node will contain the found node and idx will contain its index
	// Otherwise, it returns -1
	if idx != -1 {
		fmt.Println("Found node", node)
	} else {
		fmt.Println("Node not found")
	}
	// Output: Found node {3 [6 7]}
}

func TestFindEmpty(t *testing.T) {
	var empty []int
	_, idx := itergo.Find(empty, func(int, int) bool { return true })
	if idx != -1 {
		t.Errorf("Expected -1, got %d", idx)
	}

	_, idx = itergo.FindLast(empty, func(int, int) bool { return true })
	if idx != -1 {
		t.Errorf("Expected -1, got %d", idx)
	}
}

func TestFind(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}

	item, idx := itergo.Find(collection, func(item int, idx int) bool { return item%2 == 0 })
	if item != 2 || idx != 1 {
		t.Errorf(
			"Expected item %d at index %d, got item %d at index %d",
			2,
			1,
			item,
			idx,
		)
	}

	item, idx = itergo.FindLast(collection, func(item int, idx int) bool { return item%2 == 0 })
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
