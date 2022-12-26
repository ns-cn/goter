package goter

import "strings"

func IsYes(value string) bool {
	return strings.ToUpper(value) == "Y" || strings.ToUpper(value) == "TRUE"
}
