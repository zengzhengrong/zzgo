package zslice

// SliceRever is 切片反转
func SliceRever(items ...any) []any {
	l := len(items)
	for i := 0; i < len(items)/2; i++ {
		// 首尾 交换 len(str)/2 次
		items[i], items[l-i-1] = items[l-i-1], items[i]
	}
	return items
}
