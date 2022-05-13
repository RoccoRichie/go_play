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
