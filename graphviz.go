package itree

import (
	"fmt"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

var g = graphviz.New()

// Graphviz generates a graphviz representation of the tree
func (t Tree) Graphviz() (*cgraph.Graph, error) {
	graph, err := g.Graph()
	if err != nil {
		return nil, err
	}

	_, err = createGraphvizNode(graph, t.root)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func createGraphvizNode(graph *cgraph.Graph, in *intervalTreeNode) (*cgraph.Node, error) {
	n, err := graph.CreateNode(fmt.Sprintf("%d-%d", in.Start, in.End))
	if err != nil {
		return nil, err
	}

	if in.Interval.visited {
		n.SetFontColor("red")
	}

	if in.left != nil {
		l, err := createGraphvizNode(graph, in.left)
		if err != nil {
			return nil, err
		}
		_, err = graph.CreateEdge("links", n, l)
		if err != nil {
			return nil, err
		}
	}

	if in.right != nil {
		r, err := createGraphvizNode(graph, in.right)
		if err != nil {
			return nil, err
		}
		_, err = graph.CreateEdge("links", n, r)
		if err != nil {
			return nil, err
		}
	}

	return n, nil
}
