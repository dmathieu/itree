package itree

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntervalTreeNodeInsertLeft(t *testing.T) {
	n := newIntervalTreeNode(Interval{Start: 4, End: 6})
	err := n.insert([]Interval{Interval{Start: 1, End: 3}})
	assert.NoError(t, err)

	assert.Equal(t, 2, n.SubtreeCount)
	assert.Equal(t, int64(6), n.SubtreeMax)
	assert.Nil(t, n.right)
	assert.Equal(t, newIntervalTreeNode(Interval{Start: 1, End: 3}), n.left)
}

func TestIntervalTreeNodeInsertRight(t *testing.T) {
	n := newIntervalTreeNode(Interval{Start: 4, End: 6})
	err := n.insert([]Interval{Interval{Start: 8, End: 10}})
	assert.NoError(t, err)

	assert.Equal(t, 2, n.SubtreeCount)
	assert.Equal(t, int64(10), n.SubtreeMax)
	assert.Nil(t, n.left)
	assert.Equal(t, newIntervalTreeNode(Interval{Start: 8, End: 10}), n.right)
}

func TestIntervalTreeNodeInsertMultiple(t *testing.T) {
	n := newIntervalTreeNode(Interval{Start: 4, End: 6})
	err := n.insert([]Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 5, End: 7},
		Interval{Start: 8, End: 10},
	})
	assert.NoError(t, err)

	assert.Equal(t, &intervalTreeNode{
		Interval:     Interval{Start: 4, End: 6},
		SubtreeCount: 4,
		SubtreeMax:   10,
		left:         nil,
		right: &intervalTreeNode{
			Interval:     Interval{Start: 5, End: 7},
			SubtreeCount: 3,
			SubtreeMax:   10,
			left: &intervalTreeNode{
				Interval:     Interval{Start: 1, End: 3},
				SubtreeCount: 1,
				SubtreeMax:   3,
				left:         nil,
			},
			right: &intervalTreeNode{
				Interval:     Interval{Start: 8, End: 10},
				SubtreeCount: 1,
				SubtreeMax:   10,
				left:         nil,
				right:        nil,
			},
		},
	}, n)
}

func TestIntervalTreeNodeInsertUnbalancedLeft(t *testing.T) {
	n := newIntervalTreeNode(Interval{Start: 4, End: 6})
	err := n.insert([]Interval{Interval{Start: 1, End: 3}})
	assert.NoError(t, err)

	err = n.insert([]Interval{Interval{Start: 1, End: 3}})
	assert.Error(t, err)
}

func TestIntervalTreeNodeInsertUnbalancedRight(t *testing.T) {
	n := newIntervalTreeNode(Interval{Start: 4, End: 6})
	err := n.insert([]Interval{Interval{Start: 8, End: 10}})
	assert.NoError(t, err)

	err = n.insert([]Interval{Interval{Start: 8, End: 10}})
	assert.Error(t, err)
}

func TestIntervalTreeNodeContains(t *testing.T) {
	n := newIntervalTreeNode(Interval{Start: 1, End: 3})

	assert.False(t, n.contains(context.Background(), 5))
	assert.True(t, n.contains(context.Background(), 2))
}
