package itree

import (
	"fmt"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

var g = graphviz.New()

// GraphvizOptions allows configuring the graphviz output
type GraphvizOptions struct {
	ShowAllNodes bool
}

// Graphviz generates a graphviz representation of the tree
func (t Tree) Graphviz(opt GraphvizOptions) (*cgraph.Graph, error) {
	graph, err := g.Graph()
	if err != nil {
		return nil, err
	}

	_, err = createGraphvizNode(graph, opt, t.root)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

func createGraphvizNode(graph *cgraph.Graph, opt GraphvizOptions, in *intervalTreeNode) (*cgraph.Node, error) {
	n, err := graph.CreateNode(fmt.Sprintf("%d-%d", in.Start, in.End))
	if err != nil {
		return nil, err
	}

	if in.visited {
		n.SetFontColor("red")
	}

	if opt.ShowAllNodes || in.visited {
		if in.left != nil {
			l, err := createGraphvizNode(graph, opt, in.left)
			if err != nil {
				return nil, err
			}
			_, err = graph.CreateEdge("links", n, l)
			if err != nil {
				return nil, err
			}
		}

		if in.right != nil {
			r, err := createGraphvizNode(graph, opt, in.right)
			if err != nil {
				return nil, err
			}
			_, err = graph.CreateEdge("links", n, r)
			if err != nil {
				return nil, err
			}
		}
	}

	return n, nil
}
