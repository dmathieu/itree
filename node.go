package itree

import (
	"fmt"
)

type intervalTreeNode struct {
	Interval

	SubtreeCount int
	SubtreeMin   int64
	SubtreeMax   int64

	left  *intervalTreeNode
	right *intervalTreeNode

	visited bool
}

func newIntervalTreeNode(i Interval) *intervalTreeNode {
	return &intervalTreeNode{
		Interval:     i,
		SubtreeCount: 1,
		SubtreeMin:   i.Start,
		SubtreeMax:   i.End,
	}
}

func (tn *intervalTreeNode) insert(r []Interval) error {
	if len(r) == 0 {
		return nil
	}
	c := len(r) / 2
	e := newIntervalTreeNode(r[c])

	if e.Start > e.End {
		return fmt.Errorf("start cannot be after End for node %#v", e)
	}

	if e.End < tn.Start {
		if tn.left != nil {
			return tn.left.insert(r)
		}

		tn.left = e
	} else {
		if tn.right != nil {
			return tn.right.insert(r)
		}

		tn.right = e
	}
	err := e.insert(r[0:c])
	if err != nil {
		return err
	}
	err = e.insert(r[c+1:])
	if err != nil {
		return err
	}

	if e.SubtreeMin < tn.SubtreeMin {
		tn.SubtreeMin = e.SubtreeMin
	}
	if e.SubtreeMax > tn.SubtreeMax {
		tn.SubtreeMax = e.SubtreeMax
	}
	tn.SubtreeCount += e.SubtreeCount

	return nil
}

func (tn *intervalTreeNode) contains(value int64) bool {
	tn.visited = true

	if tn.Start <= value && tn.End >= value {
		return true
	}

	if tn.left != nil && value >= tn.left.SubtreeMin && value <= tn.left.SubtreeMax {
		return tn.left.contains(value)
	}
	if tn.right != nil && value >= tn.right.SubtreeMin && value <= tn.right.SubtreeMax {
		return tn.right.contains(value)
	}
	return false
}
