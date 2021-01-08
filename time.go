package itree

import (
	"fmt"
	"time"

	"github.com/goccy/go-graphviz/cgraph"
)

type TimesTree struct {
	Tree
}

func (t TimesTree) Contains(v time.Time) bool {
	return t.Tree.Contains(v.Unix())
}

func (t TimesTree) Graphviz(opt GraphvizOptions) (*cgraph.Graph, error) {
	opt.stringValue = func(v int64) string {
		return time.Unix(v, 0).Format(time.RFC3339)
	}

	return t.Tree.Graphviz(opt)
}

func NewTimesTree(val [][]time.Time) (TimesTree, error) {
	itvl := []Interval{}

	for _, v := range val {
		if len(v) != 2 {
			return TimesTree{}, fmt.Errorf("cannot use value %#v. Expected two values", v)
		}

		itvl = append(itvl, Interval{Start: v[0].Unix(), End: v[1].Unix()})
	}

	t, err := NewTree(itvl)
	if err != nil {
		return TimesTree{}, err
	}

	return TimesTree{t}, err
}
