package itree

import (
	"context"
)

type intervalTreeNode struct {
	Interval

	SubtreeCount int
	SubtreeMax   int64

	left  *intervalTreeNode
	right *intervalTreeNode
}

func newIntervalTreeNode(i Interval) *intervalTreeNode {
	return &intervalTreeNode{
		Interval:     i,
		SubtreeCount: 1,
		SubtreeMax:   i.End,
	}
}

func (tn *intervalTreeNode) insert(r []Interval) {
	if len(r) == 0 {
		return
	}
	c := len(r) / 2
	e := newIntervalTreeNode(r[c])

	if e.End < tn.Start {
		if tn.left == nil {
			tn.left = e
		} else {
			tn.left.insert([]Interval{r[c]})
		}
	} else {
		if tn.right == nil {
			tn.right = e
		} else {
			tn.right.insert([]Interval{r[c]})
		}
	}
	e.insert(r[0:c])
	e.insert(r[c+1:])

	if e.SubtreeMax > tn.SubtreeMax {
		tn.SubtreeMax = e.SubtreeMax
	}
	tn.SubtreeCount += e.SubtreeCount
}

func (tn *intervalTreeNode) contains(ctx context.Context, value int64) bool {
	if tn.Start <= value && tn.End >= value {
		return true
	}

	if tn.left != nil && value < tn.left.SubtreeMax {
		return tn.left.contains(ctx, value)
	}
	if tn.right != nil {
		return tn.right.contains(ctx, value)
	}
	return false
}
