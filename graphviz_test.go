package itree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphviz(t *testing.T) {
	tree, err := NewTree([]Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 5, End: 8},
		Interval{Start: 10, End: 12},
		Interval{Start: 13, End: 16},
	})
	assert.NoError(t, err)

	_ = tree.Contains(15)

	graph, err := tree.Graphviz(GraphvizOptions{})
	assert.NoError(t, err)
	assert.NotNil(t, graph)
}

func TestGraphvizShowAllNodes(t *testing.T) {
	tree, err := NewTree([]Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 5, End: 8},
		Interval{Start: 10, End: 12},
		Interval{Start: 13, End: 16},
	})
	assert.NoError(t, err)

	_ = tree.Contains(15)

	graph, err := tree.Graphviz(GraphvizOptions{ShowAllNodes: true})
	assert.NoError(t, err)
	assert.NotNil(t, graph)
}
