package command

import (
	"testing"
)

func TestPasswordSignalsRegex(t *testing.T) {
	// Real-world failure messages that should trigger the password flow.
	positives := []string{
		"sudo: a terminal is required to read the password",
		"ls: cannot open directory '/root': Permission denied",
		"sudo: a password is required",
		"Password:",
		"sudo: no askpass program specified",
	}
	for _, msg := range positives {
		if !passwordSignals.MatchString(msg) {
			t.Errorf("passwordSignals should match: %q", msg)
		}
	}

	negatives := []string{
		"command not found",
		"No such file or directory",
		"error: target not found",
	}
	for _, msg := range negatives {
		if passwordSignals.MatchString(msg) {
			t.Errorf("passwordSignals should NOT match: %q", msg)
		}
	}
}
