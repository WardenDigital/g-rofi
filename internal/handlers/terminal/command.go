package terminal

import (
	"os/exec"

	"github.com/WardenDigital/g-rofi/internal/rofi"
)

func OneShotCommand() error {
	command, err := rofi.Prompt("Enter command:")
	if err != nil {
		return err
	}

	return executeCommand(string(command))
}

func executeCommand(command string) error {
	exec.Command("sh", "-c", command).Start()
	return nil
}
