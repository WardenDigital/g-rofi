package rofi

import (
	"os"
	"os/exec"
	"strings"
)

const delimiter = "\n"

func Show(title string) ([]byte, error) {
	cmd := createShowCommand(title)

	return cmd.Output()
}

func Select(title string, options []string) ([]byte, error) {
	r, w, err := os.Pipe()

	if err != nil {
		return []byte{}, err
	}

	defer r.Close()

	piped := exec.Command("printf", createPipedOptionsString(options))
	piped.Stdout = w

	err = piped.Start()
	if err != nil {
		return []byte{}, err
	}
	defer piped.Wait()
	w.Close()

	cmd := createDmenuCommand("Select an option", r)

	return cmd.Output()
}

func createPipedOptionsString(options []string) string {
	return strings.Join(options, delimiter)
}
