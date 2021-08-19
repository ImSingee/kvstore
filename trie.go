package kvstore

import "strings"

type TrieNode struct {
	ParentKey string
	LastKey   string
	Leaf      bool
	Children  map[string]*TrieNode
}

type TrieTree struct {
	Root TrieNode
}

func (t *TrieTree) AddNew(key string) {
	parts := strings.Split(key, ".")

	node := &t.Root
	for _, part := range parts {
		node = node.EnsureGet(part)
	}
	node.Leaf = true
}

func (t *TrieTree) CheckExist(key string) bool {
	parts := strings.Split(key, ".")

	node := &t.Root
	for _, part := range parts {
		node = node.Get(part)
		if node == nil {
			return false
		}
		if node.Leaf {
			return true
		}
	}

	return false
}

func (n *TrieNode) CheckExist(key string) bool {
	if n == nil || n.Children == nil {
		return false
	}

	_, ok := n.Children[key]
	return ok
}

func (n *TrieNode) FullKey() string {
	if n.ParentKey == "" {
		return n.LastKey
	} else {
		return n.ParentKey + "." + n.LastKey
	}
}

func (n *TrieNode) Get(key string) *TrieNode {
	if n == nil || n.Children == nil {
		return nil
	}

	return n.Children[key]
}

func (n *TrieNode) EnsureGet(key string) *TrieNode {
	if n.Children == nil {
		n.Children = make(map[string]*TrieNode, 8)
	}

	if _, ok := n.Children[key]; !ok {
		n.Children[key] = &TrieNode{
			ParentKey: n.FullKey(),
			LastKey:   key,
			Children:  nil,
		}
	}

	return n.Children[key]
}

func NewTrie(names []string) *TrieTree {
	t := TrieTree{}
	for _, name := range names {
		t.AddNew(name)
	}

	return &t
}
