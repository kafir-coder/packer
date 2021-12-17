package greetings

import "fmt"

func Hello(name string) string {

	result := fmt.Sprintf("Hello %s", name)
	return result
}
