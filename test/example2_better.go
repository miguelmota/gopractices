package main

import (
	"fmt"

	// import single cohesive package
	"github.com/miguelmota/gopractices/better/example2/tree"
)

func main() {
	// more descriptive, know exactly what each constructor will provide
	myTree := tree.NewTree(tree.NewBranchNode(), tree.NewLeafNode())

	fmt.Println(myTree)
}
