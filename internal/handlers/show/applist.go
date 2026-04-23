package show

import (
	"github.com/WardenDigital/g-rofi/internal/command"
	"github.com/WardenDigital/g-rofi/internal/rofi"
)

func Show(option string) error {
	cmd, err := rofi.Show(option)
	if err != nil {
		return err
	}

	return command.ExecuteCommand(string(cmd))
}
