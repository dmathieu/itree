package itree

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimesTree(t *testing.T) {
	now := time.Now()
	tree, err := NewTimesTree([][]time.Time{
		[]time.Time{now.Add(-2 * time.Minute), now.Add(-1 * time.Minute)},
		[]time.Time{now.Add(-2 * time.Hour), now.Add(-1 * time.Hour)},
	})
	assert.NoError(t, err)
	assert.True(t, tree.Contains(now.Add(-90*time.Minute)))
	assert.False(t, tree.Contains(now))
}
