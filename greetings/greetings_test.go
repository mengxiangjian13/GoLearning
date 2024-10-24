package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "mengxiangjian"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, error := Hello(name)
	if !want.MatchString(msg) || error != nil {
		t.Fatalf(`Hello("mengxiangjian") = %q, %v, want match for %#q, nil`, msg, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
