package domiannameindex

import (
	"fmt"
	"strings"
)

const (
	wildcardNodeValue = "*"
)

type node struct {
	children map[string]*node
	value    string
}

func (n *node) insert(domainName string) {
	curr := n
	parts := strings.Split(domainName, ".")
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		child, ok := curr.children[part]
		if !ok {
			child = &node{children: make(map[string]*node), value: part}
			curr.children[part] = child
		}
		curr = child
	}
	curr.value = parts[0]
}

func (n *node) find(domainName string) (ok bool, fullPath string) {
	curr := n
	parts := strings.Split(domainName, ".")
	fullPathStackSlice := make([]string, 0, len(parts))
	for i := len(parts) - 1; i >= 0; i-- {
		fullPathStackSlice = append(fullPathStackSlice, curr.value)
		part := parts[i]
		child, ok := curr.children[part]
		if !ok {
			if _, ok := curr.children[wildcardNodeValue]; ok {
				fullPathStackSlice = append(fullPathStackSlice, wildcardNodeValue)
				Reverse[string](fullPathStackSlice)
				fullPath = strings.Join(fullPathStackSlice, ".")
				return true, fullPath[:len(fullPath)-1] // remove trailing dot
			}
			return false, ""
		}
		curr = child
	}
	fullPathStackSlice = append(fullPathStackSlice, curr.value)
	Reverse[string](fullPathStackSlice)
	fullPath = strings.Join(fullPathStackSlice, ".")
	return true, fullPath[:len(fullPath)-1] // remove trailing dot
}

func (n *node) print(prefix string) {
	fmt.Printf("%s: %-20s\n", prefix, n.value)
	for label, child := range n.children {
		child.print(fmt.Sprintf("%-20s.%s", label, prefix))
	}
}

func (n *node) remove(domainName string) {
	var stack []*node
	curr := n
	parts := strings.Split(domainName, ".")
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		child, ok := curr.children[part]
		if !ok {
			return // node not found
		}
		stack = append(stack, curr)
		curr = child
	}
	stack = append(stack, curr)
	if len(curr.children) == 0 {
		// leaf node has no children, remove it
		for len(stack) > 0 {
			parent := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(parent.children) == 0 {
				continue
			}
			label := parts[len(parts)-len(stack)-1]
			delete(parent.children, label)
			if len(parent.children) > 0 {
				break
			}

		}
	}
}

func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
