package goter

import "strings"

func isYes(value string) bool {
	return strings.ToUpper(value) == "Y" || strings.ToUpper(value) == "TRUE"
}
