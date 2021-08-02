package test

import (
	"testing"

	"github.com/zengzhengrong/zzgo/zlink"
)

func TestListNode(t *testing.T) {
	head := &zlink.LinkNode{
		Value: 1,
		Next: &zlink.LinkNode{
			Value: 2,
			Next: &zlink.LinkNode{
				Value: 3,
				Next: &zlink.LinkNode{
					Value: 4,
					Next: &zlink.LinkNode{
						Value: 5,
						Next: &zlink.LinkNode{
							Value: 6,
							Next:  nil,
						},
					},
				},
			},
		},
	}
	zlink.Check(head)
	head = head.Reverse()
	zlink.Check(head)
	head = zlink.ReverseNode(head)
	zlink.Check(head)

}
