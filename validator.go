package goter

import "os"

func Validate[T interface{}](data T, checker func(T) bool, message string, errorAction func()) {
	if checker(data) {
		_, _ = os.Stderr.WriteString("ERROR! " + message + "\n")
		if errorAction != nil {
			errorAction()
		}
		os.Exit(1)
	}
}
