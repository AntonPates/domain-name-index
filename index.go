package domiannameindex

import "sync"

const (
	rootDomainName = "."
	defaultPrefix  = "├──> "
)

// Tree is a tree of domain names.
type Tree struct {
	root  *node
	mutex sync.RWMutex
}

// New creates a new tree.
func New() Interface {
	return &Tree{root: &node{children: make(map[string]*node), value: rootDomainName}}
}

// Insert inserts a domain name into the tree.
func (t *Tree) Insert(domainName string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root.insert(domainName)
}

// Find finds a domain name in the tree and returns fullPath with respect of wildcard.
func (t *Tree) Find(domainName string) (ok bool, fullPath string) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.root.find(domainName)
}

// Remove removes a domain name from the tree.
func (t *Tree) Remove(domainName string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root.remove(domainName)
}

// Print prints the tree.
func (t *Tree) Print(prefix string) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	if len(prefix) == 0 {
		t.root.print(defaultPrefix)
		return
	}
	t.root.print(prefix)
}
