package helper

import (
	"fmt"
	"strings"
)

func LeftPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return fmt.Sprintf("%s%s", padding, s)
}

func RightPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return fmt.Sprintf("%s%s", s, padding)
}
