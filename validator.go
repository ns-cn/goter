package goter

import "os"

func Validate[T interface{}](data T, checker func(T) bool, message string) {
	if checker(data) {
		_, _ = os.Stderr.WriteString("ERROR! " + message + "\n")
		os.Exit(1)
	}
}
