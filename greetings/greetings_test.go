package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return
func TestHelloName(t *testing.T) {
	name := "Bo"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Bo")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Bo") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hllo("" = %q, %v, want "", error`, msg, err)
	}
}
