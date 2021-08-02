package zslice

import (
	"fmt"
	"strings"
)

// 任意slice合并
func SliceJoin(sep string, elems ...interface{}) string {
	l := len(elems)
	if l == 0 {
		return ""
	}
	if l == 1 {
		s := fmt.Sprint(elems[0])
		sLen := len(s) - 1
		if s[0] == '[' && s[sLen] == ']' {
			return strings.Replace(s[1:sLen], " ", sep, -1)
		}
		return s
	}
	sep = strings.Replace(fmt.Sprint(elems), " ", sep, -1)
	return sep[1 : len(sep)-1]
}
