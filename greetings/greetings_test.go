package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName tests the behavior of the Hello function when called with a name.
// It checks if the function returns a valid greeting message.
func TestHelloName(t *testing.T) {
	// Set the name for testing.
	name := "Gladys"
	// Create a regular expression pattern to match the expected greeting.
	want := regexp.MustCompile(`\b` + name + `\b`)
	// Call the Hello function with the given name.
	msg, err := Hello("Gladys")
	// Check if the greeting message matches the expected pattern and if there is no error.
	if !want.MatchString(msg) || err != nil {
		// Log a test failure if the condition is not satisfied.
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty tests the behavior of the Hello function when called with an empty string.
// It checks if the function returns an error, indicating that an empty name is not allowed.
func TestHelloEmpty(t *testing.T) {
	// Call the Hello function with an empty string.
	msg, err := Hello("")
	// Check if the greeting message is not empty and if there is an error.
	if msg != "" || err == nil {
		// Log a test failure if the condition is not satisfied.
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
