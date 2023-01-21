package index

import (
	"fmt"

	"github.com/miekg/dns"
)

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{Root: New()}
}

func (t *Tree) Add(domainName string) {
	domainName = normalize(domainName)
	labels := dns.SplitDomainName(domainName)
	reverse(labels)
	// Start at the root node.
	node := t.Root
	for _, label := range labels {
		// Check if the node has a child with the current component.
		if ok, wildcard := node.HasChild(label); ok && !wildcard {
			// Get the child node.
			node = node.GetChild(label)
		} else {
			// Add a new child node.
			node = node.AddChild(label)
		}
	}
}

func (t *Tree) Remove(domainName string) {
	domainName = normalize(domainName)
	labels := dns.SplitDomainName(domainName)
	reverse(labels)
	// Split the domain name into its components.
	// Start at the root node.
	node := t.Root
	// Iterate over the components.
	for i, label := range labels {
		// Check if the node has a child with the current component.
		if ok, _ := node.HasChild(label); ok {
			// Get the child node.
			fmt.Println("label", label)
			node = node.GetChild(label)
			if i == len(labels)-1 && node.AmountOfChildren() == 0 {
				node.Remove()
			}
		} else {
			// The domain name does not exist in the index.
			return
		}
	}
}

func (t *Tree) Has(domainName string) (has bool, isWildcard bool) {
	domainName = normalize(domainName)
	labels := dns.SplitDomainName(domainName)
	reverse(labels)
	// Start at the root node.
	node := t.Root
	// Iterate over the components.
	for _, label := range labels {
		if ok, wildcard := node.HasChild(label); ok {
			if wildcard {
				return true, true
			}
			node = node.GetChild(label)
		} else {
			// The domain name does not exist in the index.
			return false, false
		}
	}
	// The domain name exists in the index.
	return true, false
}

// Reverse the components.
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
