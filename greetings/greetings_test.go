package greetings

import (
	"regexp"
	"testing"
)

// calls greetings.Hello with a name, checking for a valid return value.
func TestHelloName(t *testing.T) {
	name := "trogdor"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("trogdor")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("trogdor") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// this one calls with an empty string, checking for error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "" error`, msg, err)
	}
}
