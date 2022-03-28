package test

import (
	"fmt"
	"testing"

	"github.com/zengzhengrong/zzgo/zslice"
)

func TestMergeSort(t *testing.T) {

	// list1 := []int{1, 2, 3, 4}
	list2 := []float32{5.1, 4, 2, 3, 3, 2, 1}
	result := zslice.MergeSort(list2)
	fmt.Println(result)
	// for len(list1) > 0 && len(list2) > 0 {
	// 	fmt.Println(list1)
	// 	// list1 = append(list1[:0], list1[0+1:]...)
	// 	copy(list1[0:], list1[0+1:])
	// 	list1[len(list1)-1] = 0 // or the zero value of T
	// 	list1 = list1[:len(list1)-1]
	// 	fmt.Println(list1)
	// }
}
