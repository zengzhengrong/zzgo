package zslice

import (
	"math"
)

//SliceSplitWithCallBack is 切片分割, 按num 个数进行分割,使用回调函数执行分割批次
func SliceSplitWithCallBack[T any](num int, list []T, callback func(batch []T)) error {
	for i := 1; i <= int(math.Floor(float64(len(list)/num)))+1; i++ {
		// 将list 除于 num 获取 分割次数 非整数时 舍去小数点 后+1
		low := num * (i - 1) // 左索引

		if low > len(list) {
			// 左索引 大于长度列表长度则直接返回
			return nil
		}
		high := num * i // 右索引
		if high > len(list) {
			// 如果右索引大于 list长度，则取list的长度
			high = len(list)
		}
		callback(list[low:high])
	}
	return nil
}

//SliceSplitWithChannel is 切片分割, 按num 个数进行分割,返回channel
func SliceSplitWithChannel[T any](num int, list []T) <-chan []T {

	// v := reflect.ValueOf(list)
	batchSize := int(math.Floor(float64(len(list)/num))) + 1
	c := make(chan []T, batchSize)
	defer close(c)

	for i := 1; i <= batchSize; i++ {

		// 将list 除于 num 获取 分割次数 非整数时 舍去小数点 后+1
		low := num * (i - 1) // 左索引

		if low > len(list) {
			// 左索引 大于长度列表长度则直接返回
			close(c)
		}
		high := num * i // 右索引
		if high > len(list) {
			// 如果右索引大于 list长度，则取list的长度
			high = len(list)
		}
		// c <- v.Slice(low, high).Interface().(T)
		c <- list[low:high]
	}

	return c
}
