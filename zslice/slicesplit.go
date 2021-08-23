package zslice

import (
	"math"
)

//SliceSplit is 切片分割, 按num 个数进行分割
func SliceSplit(num int, list []interface{}, exec func(l []interface{}, low, hight int)) {

	for i := 1; i <= int(math.Floor(float64(len(list)/num)))+1; i++ {
		// 将list 除于 num 获取 分割次数 非整数时 舍去小数点 后+1
		low := num * (i - 1) // 左索引

		if low > len(list) {
			// 左索引 大于长度列表长度则直接返回
			return
		}
		high := num * i // 右索引
		if high > len(list) {
			// 如果右索引大于 list长度，则取list的长度
			high = len(list)
		}
		exec(list, low, high)
	}
}
