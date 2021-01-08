package itree

import (
	"encoding/binary"
	"net"

	"github.com/goccy/go-graphviz/cgraph"
	"gopkg.in/netaddr.v1"
)

type IPNetTree struct {
	Tree
}

func (t IPNetTree) Contains(v net.IP) bool {
	return t.Tree.Contains(int64(ipV4ToInt(v)))
}

func (t IPNetTree) Graphviz(opt GraphvizOptions) (*cgraph.Graph, error) {
	opt.stringValue = func(v int64) string {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, uint32(v))
		return ip.String()
	}

	return t.Tree.Graphviz(opt)
}

func NewIPNetTree(val []*net.IPNet) (IPNetTree, error) {
	itvl := []Interval{}

	for _, v := range val {
		start := int64(ipV4ToInt(v.IP))
		end := int64(ipV4ToInt(netaddr.BroadcastAddr(v)))

		itvl = append(itvl, Interval{Start: start, End: end})
	}

	t, err := NewTree(itvl)
	if err != nil {
		return IPNetTree{}, err
	}

	return IPNetTree{t}, err
}

func ipV4ToInt(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}
