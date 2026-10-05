package command

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/WardenDigital/g-rofi/internal/rofi"
)

const maxPasswordAttempts = 3

// passwordSignals matches stderr output that suggests a command failed
// because it needs a password or elevated privileges (e.g. sudo run without
// a tty, or a permission error).
var passwordSignals = regexp.MustCompile(`(?i)password|permission denied|no tty|not a tty|terminal is required|askpass`)

// ExecuteCommand runs a command through the shell. If the first attempt
// fails and the error output suggests a password is required, the user is
// prompted for a password (masked, via rofi) and the command is retried
// through sudo -S.
func ExecuteCommand(command string) error {
	stderr, err := runShellCommand(command, "")
	if err != nil && passwordSignals.MatchString(stderr) {
		return runWithPassword(command)
	}

	return err
}

// runShellCommand executes command via sh. When password is non-empty the
// command is run with sudo -S and the password is fed to it on stdin.
// Command stderr is both printed to the user and captured for inspection.
func runShellCommand(command, password string) (string, error) {
	var stderr bytes.Buffer

	cmd := exec.Command("sh", "-c", command)

	if password != "" {
		cmd = exec.Command("sudo", "-p", "", "-S", "sh", "-c", command)
		cmd.Stdin = strings.NewReader(password + "\n")
	}

	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)
	err := cmd.Run()

	return stderr.String(), err
}

// runWithPassword prompts for a password and retries the command with
// sudo -S until it succeeds, the user cancels, or the failure no longer
// looks password-related.
func runWithPassword(command string) error {
	var lastErr error

	for range maxPasswordAttempts {
		password, err := rofi.Password("Password required:")
		if err != nil {
			return err
		}

		stderr, err := runShellCommand(command, string(password))
		if err == nil {
			return nil
		}

		lastErr = err

		if !passwordSignals.MatchString(stderr) {
			return err
		}
	}

	return lastErr
}