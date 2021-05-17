package log

import "fmt"

var dev = false

func Info(content ...interface{}) {
	fmt.Print("info", " ", content, "\n")
}

func Error(content ...interface{}) {
	fmt.Print("error", " ", content, "\n")
}

func Dev(content ...interface{}) {
	if dev {
		fmt.Print("error", " ", content, "\n")
	}
}

// @dev Gives error some context msg.
func WrapErr(err error, msg string) (errWithContext error) {
	if err == nil {
		return
	}

	errWithContext = fmt.Errorf("%s: %v", msg, err)
	return
}
