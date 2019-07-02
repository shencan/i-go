package utils

import (
	"fmt"
	"regexp"
)

// 电话校验
func PhoneVerify(phone string) bool {
	matched, err := regexp.MatchString("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$", phone)
	if err != nil {
		fmt.Println(err)
	}
	return matched
}