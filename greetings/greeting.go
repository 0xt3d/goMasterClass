package greetings

import (
	"fmt"

	"golang.org/x/text/message"
)

func Hello(name string) string {
	message := fmt.Sprint("Hi, %v. Welcome!, name")
	return message
}