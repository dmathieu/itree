package itree

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestFuzz(t *testing.T) {
	f := fuzz.New().NilChance(0).Funcs(func(i *Interval, c fuzz.Continue) {
		start := c.Int63()
		end := c.Int63()
		if start > end {
			tmp := end
			end = start
			start = tmp
		}

		i.Start = start
		i.End = end
	})

	var i []Interval
	f.Fuzz(&i)

	tree, err := NewTree(i)
	assert.NoError(t, err)

	for i := 0; i <= 1000; i++ {
		var i int64
		f.Fuzz(&i)
		_ = tree.Contains(i)
	}
}
