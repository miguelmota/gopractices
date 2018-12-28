package main

import (
	"fmt"

	// need to import separate package for each type even though they are all related
	"github.com/miguelmota/gopractices/bad/example2/branchnode"
	"github.com/miguelmota/gopractices/bad/example2/leafnode"
	"github.com/miguelmota/gopractices/bad/example2/tree"
)

func main() {
	treeSvc := tree.New(branchnode.New(), leafnode.New())

	// also, it doesn't make sense to have `tree.Service`, `branchnode.Service`, `leafnode.Service` because they are simple data structures and not part of a service oriented architecture

	fmt.Println(treeSvc)
}
