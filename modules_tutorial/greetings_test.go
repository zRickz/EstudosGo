package greetings

import (
	"regexp"
	"testing"
)

// Teste para Hello com nome
func TestHelloName(t *testing.T) {
	want := regexp.MustCompile(`\b` + "Joãozinho" + `\b`)
	msg, err := Hello("Joãozinho")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Joãozinho") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Teste para Hello sem nome/vázio
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want "", error`, msg, err)
	}
}
