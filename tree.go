package itree

import (
	"context"
	"sort"
)

type Tree struct {
	root *intervalTreeNode
}

func NewTree(itvl []Interval) (Tree, error) {
	var tree Tree

	if len(itvl) == 0 {
		return tree, nil
	}

	sort.Slice(itvl, func(i, j int) bool {
		return itvl[i].End > itvl[j].End
	})

	rID := len(itvl) / 2
	tree.root = newIntervalTreeNode(itvl[rID])
	err := tree.root.insert(itvl[0:rID])
	if err != nil {
		return tree, err
	}
	err = tree.root.insert(itvl[rID+1:])
	if err != nil {
		return tree, err
	}

	return tree, nil
}

func (t Tree) Contains(ctx context.Context, value int64) bool {
	if t.root == nil {
		return false
	}

	return t.root.contains(ctx, value)
}
