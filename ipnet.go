package itree

import (
	"encoding/binary"
	"net"

	"gopkg.in/netaddr.v1"
)

type IPNetTree struct {
	Tree
}

func (t IPNetTree) Contains(v net.IP) bool {
	return t.Tree.Contains(int64(ipV4ToInt(v)))
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
