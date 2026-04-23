package command

import "os/exec"

func ExecuteCommand(command string) error {
	exec.Command("sh", "-c", command).Start()
	return nil
}
