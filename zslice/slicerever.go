package zslice

// SliceRever is 切片反转
func SliceRever(list []interface{}) []interface{} {
	l := len(list)
	for i := 0; i < len(list)/2; i++ {
		// 首尾 交换 len(str)/2 次
		list[i], list[l-i-1] = list[l-i-1], list[i]
	}
	return list
}
