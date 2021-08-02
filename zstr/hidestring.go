package zstr

//HideString is 隐藏中间字符用*号代替
func HideString(src string, hLen int) string {
	str := []rune(src)
	if hLen == 0 {
		hLen = 4
	}
	hideStr := ""
	for i := 0; i < hLen; i++ {
		hideStr += "*"
	}
	hideLen := len(str) / 2
	showLen := len(str) - hideLen
	if hideLen == 0 || showLen == 0 {
		return hideStr
	}
	subLen := showLen / 2
	if subLen == 0 {
		return string(str[:showLen]) + hideStr
	}
	s := string(str[:subLen])
	s += hideStr
	s += string(str[len(str)-subLen:])
	return s
}
