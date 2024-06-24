package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Omer"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Omer")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Omer") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
