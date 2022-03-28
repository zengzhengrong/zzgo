package zslice

import (
	"constraints"
	"fmt"
)

//[]int{5, 4, 2, 3, 3, 2, 1}
func MergeSort[T constraints.Ordered](list []T) []T {
	if len(list) == 1 {
		return list
	}

	left, right := split(list)
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

func split[T constraints.Ordered](list []T) ([]T, []T) {
	mid := len(list) / 2
	return list[:mid], list[mid:]
}

func merge[T constraints.Ordered](left, right []T) []T {
	fmt.Println(left, right)
	fmt.Println(len(left) + len(right))
	result := make([]T, 0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = append(right[:0], right[0+1:]...)

		} else {
			result = append(result, left[0])
			left = append(left[:0], left[0+1:]...)
		}
	}

	for len(left) > 0 && len(right) == 0 {
		// 处理左边有元素右边没有元素
		result = append(result, left[0])
		left = append(left[:0], left[0+1:]...)
	}

	for len(right) > 0 && len(left) == 0 {
		result = append(result, right[0])
		right = append(right[:0], right[0+1:]...)
	}

	return result

}
