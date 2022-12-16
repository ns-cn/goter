package goter

import "os"

func Validate[T interface{}](data T, checker func(T) bool, message string) {
	if !checker(data) {
		_, _ = os.Stderr.WriteString(message)
		os.Exit(1)
	}
}
