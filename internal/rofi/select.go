package rofi

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const delimiter = "\n"

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
	fmt.Println("Piped options string created")

	cmd := createDmenuCommand("Select an option", r)
	fmt.Println("Rofi command created")

	return cmd.Output()
}

func createPipedOptionsString(options []string) string {
	var piped strings.Builder
	for _, option := range options {
		piped.WriteString(option + delimiter)
	}

	return piped.String()
}
