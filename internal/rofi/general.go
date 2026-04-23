package rofi

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

var configFileName = "config.rasi"
var inlineTheme string = "inlinePrompt.rasi"

func createDmenuCommand(title string, r io.Reader) *exec.Cmd {
	cmd := exec.Command("rofi", "-config", getConfig(), "-dmenu", "-p", title)
	cmd.Stdin = r

	return cmd
}

func createShowCommand(show string) *exec.Cmd {
	cmd := exec.Command("rofi", "-config", getConfig(), "-show", show)

	return cmd
}

func createInlineDmenuCommand(title string, r io.Reader) *exec.Cmd {
	cmd := exec.Command("rofi", "-config", getConfig(), "-dmenu", "-p", title, "-theme", getTheme(inlineTheme))
	cmd.Stdin = r

	return cmd
}

func getConfigPath() string {
	if path := os.Getenv("G_ROFI_CONFIG"); path != "" {
		return path
	}

	if path := os.Getenv("THEME_DIR"); path != "" {
		return path
	}

	return "."
}

func getConfig() string {
	return fmt.Sprintf("%s/%s", getConfigPath(), configFileName)
}

func getTheme(name string) string {
	return fmt.Sprintf("%s/%s", getConfigPath(), name)
}
