package index

import "testing"

func TestIndex(t *testing.T) {
	// Create a new index.
	idx := New()

	// Add a new node to the index.
	node := idx.AddChild("foo")

	// Add a child to the node.
	node.AddChild("bar")

	// Add a child to the node.
	node.AddChild("baz")

	// Add a child to the node.
	node.AddChild("qux")

	// Remove the "bar" child from the node.
	node.RemoveChild("bar")

	// Remove the "baz" child from the node.
	node.RemoveChild("baz")

	// Remove the "qux" child from the node.
	node.RemoveChild("qux")

	// Remove the node from the index.
	node.Remove()

	// Check that the index has no children.
	if len(idx.Children) != 0 {
		t.Fatalf("expected index to have no children")
	}
}
