package zdup

// RemoveDuplicateString is 字符串去重
func RemoveDuplicateString(items []string) []string {
	res := make([]string, 0, len(items))
	temp := map[string]struct{}{}
	for _, item := range items {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}

// RemoveDuplicateInt is int去重
func RemoveDuplicateInt(elms []int) []int {
	res := make([]int, 0, len(elms))
	temp := map[int]struct{}{}
	for _, item := range elms {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}

// RemoveDuplicateInt64 is int64去重
func RemoveDuplicateInt64(elms []int64) []int64 {
	res := make([]int64, 0, len(elms))
	temp := map[int64]struct{}{}
	for _, item := range elms {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}

// HasDuplicateInt64 is int64去重
func HasDuplicateInt64(elms []int64) bool {
	temp := map[int64]struct{}{}
	for _, item := range elms {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
		} else {
			return true
		}

	}
	return false
}

// DuplicateInt64 is 含有重复元素
func DuplicateInt64(elms []int64) []int64 {
	result := make([]int64, 0)
	temp := map[int64]struct{}{}
	for _, item := range elms {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
		} else {
			result = append(result, item)
		}

	}
	return result
}
