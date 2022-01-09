package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zslice"
)

func TestSliceSplit(t *testing.T) {
	var list []int
	for i := 0; i < 10003; i++ {
		list = append(list, i)
	}
	callback := func(batch any) {
		list := batch.([]int)
		fmt.Println(list)
	}
	zslice.SliceSplitWithCallBack(1000, list, callback)
}
func BenchmarkSliceSplit(b *testing.B) {
	var list []int
	for i := 0; i < 10003; i++ {
		list = append(list, i)
	}
	callback := func(batch any) {
		list := batch.([]int)
		fmt.Println(list)

	}
	for n := 0; n < b.N; n++ {
		zslice.SliceSplitWithCallBack(1000, list, callback)
	}

}
func BenchmarkSliceSplitChan(b *testing.B) {
	var list []int
	for i := 0; i < 10003; i++ {
		list = append(list, i)
	}
	for n := 0; n < b.N; n++ {
		for batch := range zslice.SliceSplitWithChannel(1000, list) {
			list := batch.([]int)
			fmt.Println(list)
		}
	}
}
