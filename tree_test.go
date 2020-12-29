package itree

import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
	tree, err := NewTree([]Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 5, End: 8},
		Interval{Start: 10, End: 12},
		Interval{Start: 13, End: 16},
	})

	assert.NoError(t, err)
	assert.Equal(t, 4, tree.root.SubtreeCount)
	assert.Equal(t, int64(16), tree.root.SubtreeMax)
}

func TestTreeContains(t *testing.T) {
	tree, err := NewTree([]Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 5, End: 8},
		Interval{Start: 10, End: 12},
		Interval{Start: 13, End: 16},
	})

	assert.NoError(t, err)
	assert.True(t, tree.Contains(context.Background(), 5))
	assert.False(t, tree.Contains(context.Background(), 9))
}

func TestTreeContainsOverlap(t *testing.T) {
	tree, err := NewTree([]Interval{
		Interval{Start: 1, End: 3},
		Interval{Start: 5, End: 8},
		Interval{Start: 7, End: 12},
		Interval{Start: 13, End: 16},
	})

	assert.NoError(t, err)
	assert.True(t, tree.Contains(context.Background(), 9))
	assert.False(t, tree.Contains(context.Background(), 4))
}

func BenchmarkBuildTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intervals := []Interval{}
		for j := 0; j < b.N; j++ {
			first := rand.Int63()
			second := first + int64(rand.Intn(100))
			intervals = append(intervals, Interval{Start: first, End: second})
		}

		_, err := NewTree(intervals)
		assert.NoError(b, err)
	}
}

func BenchmarkTreeContains(b *testing.B) {
	intervals := []Interval{}
	for i := 0; i < b.N; i++ {
		first := rand.Int63()
		second := first + int64(rand.Intn(100))
		intervals = append(intervals, Interval{Start: first, End: second})
	}

	tree, err := NewTree(intervals)
	assert.NoError(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Contains(context.Background(), rand.Int63())
	}
}
