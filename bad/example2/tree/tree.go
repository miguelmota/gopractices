package tree

import (
	"github.com/miguelmota/gopractices/bad/example2/node"
)

// Service ...
type Service struct {
}

// New ...
func New(nodes ...node.Node) *Service {
	return &Service{}
}
