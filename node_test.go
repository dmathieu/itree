package itree

import (
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
		SubtreeMin:   1,
		SubtreeMax:   10,
		left:         nil,
		right: &intervalTreeNode{
			Interval:     Interval{Start: 5, End: 7},
			SubtreeCount: 3,
			SubtreeMin:   1,
			SubtreeMax:   10,
			left: &intervalTreeNode{
				Interval:     Interval{Start: 1, End: 3},
				SubtreeCount: 1,
				SubtreeMin:   1,
				SubtreeMax:   3,
				left:         nil,
			},
			right: &intervalTreeNode{
				Interval:     Interval{Start: 8, End: 10},
				SubtreeCount: 1,
				SubtreeMin:   8,
				SubtreeMax:   10,
				left:         nil,
				right:        nil,
			},
		},
	}, n)
}

func TestIntervalTreeNodeInsertChildLeft(t *testing.T) {
	n := newIntervalTreeNode(Interval{Start: 8, End: 9})
	err := n.insert([]Interval{Interval{Start: 5, End: 7}})
	assert.NoError(t, err)

	err = n.insert([]Interval{Interval{Start: 2, End: 4}})
	assert.NoError(t, err)
	assert.True(t, n.contains(3))
}

func TestIntervalTreeNodeInsertChildRight(t *testing.T) {
	n := newIntervalTreeNode(Interval{Start: 4, End: 6})
	err := n.insert([]Interval{Interval{Start: 8, End: 10}})
	assert.NoError(t, err)

	err = n.insert([]Interval{Interval{Start: 11, End: 13}})
	assert.NoError(t, err)
	assert.True(t, n.contains(12))
}

func TestIntervalTreeNodeContains(t *testing.T) {
	n := newIntervalTreeNode(Interval{Start: 1, End: 3})

	assert.False(t, n.contains(5))
	assert.True(t, n.contains(2))
}
