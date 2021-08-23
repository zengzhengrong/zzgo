package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zslice"
)

func TestSliceSplit(t *testing.T) {
	var list []interface{}
	for i := 0; i < 10003; i++ {
		list = append(list, i)
	}
	f := func(l []interface{}, low, hight int) {
		fmt.Println(l[low:hight])
	}
	zslice.SliceSplit(1000, list, f)
}
