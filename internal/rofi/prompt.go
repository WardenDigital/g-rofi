package rofi

import (
	"os"
)

func Prompt(title string) ([]byte, error) {

	cmd := createInlineDmenuCommand(title, os.Stdin)

	return cmd.Output()
}

// Password prompts for a password with the input masked in rofi.
func Password(title string) ([]byte, error) {
	cmd := createPasswordDmenuCommand(title, os.Stdin)

	return cmd.Output()
}
