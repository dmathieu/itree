package itree

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTreeFromIPNet(t *testing.T) {
	ip, ipnet1, err := net.ParseCIDR("127.0.0.1/32")
	assert.NoError(t, err)
	_, ipnet2, err := net.ParseCIDR("192.168.0.1/8")
	assert.NoError(t, err)

	tree, err := NewIPNetTree([]*net.IPNet{ipnet1, ipnet2})
	assert.NoError(t, err)
	assert.True(t, tree.Contains(ip))
	assert.False(t, tree.Contains(net.ParseIP("8.8.8.8")))
}
