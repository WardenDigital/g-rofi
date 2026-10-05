package command

import (
	"bytes"
	"errors"
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
var passwordSignals = regexp.MustCompile(`(?i)password|permission denied|no tty|not a tty|terminal is required|askpass|incorrect password|sorry, try again`)

// errPasswordDeclined is returned when the user dismisses the password
// prompt or submits an empty password.
var errPasswordDeclined = errors.New("password prompt dismissed")

// ExecuteCommand runs a command through the shell, elevating to root via
// sudo when necessary. The privilege check happens up front with `sudo -n`:
// with cached credentials the command runs immediately, and without them
// sudo fails instantly with a password-required error — so the password
// prompt shows up right away instead of only after the unprivileged run
// has tripped over missing privileges (which can take minutes for polkit
// or sudo-style commands).
func ExecuteCommand(command string) error {
	stderr, err := runWithSudo(command, "")
	if err == nil {
		return nil
	}

	// sudo not installed: just run the command as the current user.
	if errors.Is(err, exec.ErrNotFound) {
		_, err = runAsUser(command)
		return err
	}

	// sudo refused to even start the command (no cached credentials,
	// requiretty, etc.): prompt for a password and retry elevated.
	if passwordSignals.MatchString(stderr) {
		err = runWithPassword(command)
		if errors.Is(err, errPasswordDeclined) {
			// User dismissed the prompt — fall back to running as the
			// current user, which still works for non-privileged commands.
			_, err = runAsUser(command)
		}
		return err
	}

	// The command failed (or sudo refused) for a reason unrelated to
	// passwords — retry as the current user, like the original flow did.
	_, err = runAsUser(command)
	return err
}

// runAsUser executes command via sh as the current user. stderr is both
// printed to the user and captured for inspection.
func runAsUser(command string) (string, error) {
	return execute(exec.Command("sh", "-c", command))
}

// runWithSudo executes command via sudo. With an empty password it runs
// non-interactively (`sudo -n`) and fails fast when a password would be
// required. Otherwise the password is fed to sudo via stdin (`sudo -S`).
func runWithSudo(command, password string) (string, error) {
	var cmd *exec.Cmd
	if password == "" {
		cmd = exec.Command("sudo", "-n", "sh", "-c", command)
	} else {
		cmd = exec.Command("sudo", "-p", "", "-S", "sh", "-c", command)
		cmd.Stdin = strings.NewReader(password + "\n")
	}

	return execute(cmd)
}

// execute runs cmd, tee-ing stderr to the user while keeping a copy.
func execute(cmd *exec.Cmd) (string, error) {
	var stderr bytes.Buffer

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
			return errPasswordDeclined
		}

		if len(strings.TrimSpace(string(password))) == 0 {
			return errPasswordDeclined
		}

		stderr, err := runWithSudo(command, string(password))
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
