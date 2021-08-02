package zvalid

import (
	"fmt"
	"regexp"
)

// ValidatePhoneNumber is 校验手机号
func ValidatePhoneNumber(number string) (bool, error) {
	pattern := "^1[3-9]\\d{9}$" //反斜杠要转义
	result, _ := regexp.MatchString(pattern, number)
	if result {
		return result, nil
	}
	return result, fmt.Errorf("%s phone number format error", number)
}
