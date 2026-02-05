package rofi

import (
	"io"
	"os/exec"
)

var inlineTheme string = "inlinePrompt.rasi"

func createDmenuCommand(title string, r io.Reader) *exec.Cmd {
	cmd := exec.Command("rofi", "-dmenu", "-p", title)
	cmd.Stdin = r

	return cmd
}

func createInlineDmenuCommand(title string, r io.Reader) *exec.Cmd {
	cmd := exec.Command("rofi", "-dmenu", "-p", title, "-theme", inlineTheme)
	cmd.Stdin = r

	return cmd
}
