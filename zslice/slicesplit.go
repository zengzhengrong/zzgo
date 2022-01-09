package zslice

import (
	"errors"
	"math"
	"reflect"
)

//SliceSplitWithCallBack is 切片分割, 按num 个数进行分割,使用回调函数执行分割批次
func SliceSplitWithCallBack(num int, list any, callback func(batch any)) error {
	v := reflect.ValueOf(list)
	kind := v.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return errors.New("list is not slice or array")
	}

	for i := 1; i <= int(math.Floor(float64(v.Len()/num)))+1; i++ {
		// 将list 除于 num 获取 分割次数 非整数时 舍去小数点 后+1
		low := num * (i - 1) // 左索引

		if low > v.Len() {
			// 左索引 大于长度列表长度则直接返回
			return nil
		}
		high := num * i // 右索引
		if high > v.Len() {
			// 如果右索引大于 list长度，则取list的长度
			high = v.Len()
		}

		callback(v.Slice(low, high).Interface())
	}
	return nil
}

//SliceSplitWithChannel is 切片分割, 按num 个数进行分割,返回channel
func SliceSplitWithChannel(num int, list any) <-chan any {

	v := reflect.ValueOf(list)
	batchSize := int(math.Floor(float64(v.Len()/num))) + 1
	c := make(chan any, batchSize)
	defer close(c)

	for i := 1; i <= batchSize; i++ {

		// 将list 除于 num 获取 分割次数 非整数时 舍去小数点 后+1
		low := num * (i - 1) // 左索引

		if low > v.Len() {
			// 左索引 大于长度列表长度则直接返回
			close(c)
		}
		high := num * i // 右索引
		if high > v.Len() {
			// 如果右索引大于 list长度，则取list的长度
			high = v.Len()
		}
		c <- v.Slice(low, high).Interface()
	}

	return c
}
