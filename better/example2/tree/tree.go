package tree

// Node ...
type Node interface{}

// BranchNode ...
type BranchNode struct{}

// LeafNode ...
type LeafNode struct{}

// Tree ...
type Tree struct{}

// NewBranchNode ...
func NewBranchNode() Node {
	return BranchNode{}
}

// NewLeafNode ...
func NewLeafNode() Node {
	return LeafNode{}
}

// NewTree ...
func NewTree(nodes ...Node) *Tree {
	return &Tree{}
}
