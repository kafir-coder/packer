package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {

	name := "Gladys"
	want := regexp.MustCompile(`\b` + name)
	msg := Hello("Gladys")

	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q, want match for %#q, nil`, msg, want)
	}
}
