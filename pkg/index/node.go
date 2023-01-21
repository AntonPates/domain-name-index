package index

const (
	WildcardNodeValue = "*"
)

type Node struct {
	Value    string
	Children map[string]*Node
	Parent   *Node
}

func (n *Node) HasChild(value string) (has bool, isWildcard bool) {
	if n.Children == nil {
		return false, false
	}
	if _, ok := n.Children[value]; ok {
		return true, value == WildcardNodeValue
	}

	if _, ok := n.Children[WildcardNodeValue]; ok {
		return true, true
	}
	return false, false
}

func (n *Node) GetChild(value string) *Node {
	if n.Children == nil {
		return nil
	}
	return n.Children[value]
}

func (n *Node) AddChild(value string) *Node {
	child := &Node{Value: value, Parent: n}
	if n.Children == nil {
		n.Children = make(map[string]*Node)
	}
	n.Children[value] = child
	return child
}

func (n *Node) RemoveChild(value string) int {
	if n.Children == nil {
		return 0
	}
	delete(n.Children, value)
	return n.AmountOfChildren()
}

func (n *Node) AmountOfChildren() int {
	return len(n.Children)
}

func (n *Node) Remove() {
	if n.AmountOfChildren() > 0 {
		return
	}
	if n.Parent != nil {
		amountRemainedChildren := n.Parent.RemoveChild(n.Value)
		for amountRemainedChildren == 0 {
			n = n.Parent
			if n.Value == Root {
				break
			}
			amountRemainedChildren = n.Parent.RemoveChild(n.Value)
		}
	}
}
