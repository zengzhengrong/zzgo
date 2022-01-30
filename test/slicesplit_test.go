package test

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
	"github.com/zengzhengrong/zzgo/zslice"
)

func TestSliceSplit(t *testing.T) {
	var list []string
	for i := 0; i < 10003; i++ {
		list = append(list, cast.ToString(i))
	}
	callback := func(batch []string) {
		fmt.Println(batch)
	}
	zslice.SliceSplitWithCallBack(1000, list, callback)
}
func BenchmarkSliceSplit(b *testing.B) {
	var list []int
	for i := 0; i < 10003; i++ {
		list = append(list, i)
	}
	callback := func(batch []int) {
		fmt.Println(batch)

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
			fmt.Println(batch)
		}
	}
}
