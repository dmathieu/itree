package itree

import (
	"bytes"
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

	graph, err := Graphviz(tree)
	assert.NoError(t, err)

	var buf bytes.Buffer
	err = g.Render(graph, "dot", &buf)
	assert.NoError(t, err)

	assert.Equal(t, `digraph "" {
	graph [bb="0,0,139.57,180"];
	node [fontcolor=black,
		label="\N"
	];
	"5-8"	 [fontcolor=red,
		height=0.5,
		pos="66,162",
		width=0.75];
	"1-3"	 [height=0.5,
		pos="27,90",
		width=0.75];
	"5-8" -> "1-3" [key=links,
	pos="e,36.176,106.94 56.758,144.94 52.046,136.24 46.218,125.48 40.972,115.79"];
"10-12" [fontcolor=red,
	height=0.5,
	pos="106,90",
	width=0.93237];
"5-8" -> "10-12" [key=links,
pos="e,96.311,107.44 75.479,144.94 80.241,136.37 86.116,125.79 91.435,116.22"];
"13-16" [fontcolor=red,
height=0.5,
pos="106,18",
width=0.93237];
"10-12" -> "13-16" [key=links,
pos="e,106,36.413 106,71.831 106,64.131 106,54.974 106,46.417"];
}
`, buf.String())
}
