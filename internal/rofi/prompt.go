package rofi

import (
	"os"
)

func Prompt(title string) ([]byte, error) {

	cmd := createInlineDmenuCommand(title, os.Stdin)

	return cmd.Output()
}
