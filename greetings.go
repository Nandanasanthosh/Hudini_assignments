package greetings

import (
	"fmt"
)

func Hello(name string) string {
	message := fmt.Sprintf("H, %v. Welcome!", name)
	return message
}
