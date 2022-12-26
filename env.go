package goter

import "strings"

// IsYes 判断环境变量的值是否是预期的逻辑是的值
func IsYes(value string) bool {
	return strings.ToUpper(value) == "Y" || strings.ToUpper(value) == "TRUE"
}
