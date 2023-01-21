package index

const (
	// Root is the root node of the index tree.
	Root = "."
)

func New() *Node {
	return &Node{Value: Root}
}
