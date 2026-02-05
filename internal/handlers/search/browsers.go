package search

import (
	"os/exec"
	"strings"
)

type Browser interface {
	Open(url string) error
	OpenInNewWindow(url string) error
	OpenInIncognito(url string) error
}

type BraveBrowser struct {
	binaryPath string
}

func (b *BraveBrowser) Open(url string) error {
	exec.Command(b.binaryPath, url).Start()

	return nil
}

func (b *BraveBrowser) OpenInNewWindow(url string) error {
	exec.Command("brave-browser", "--new-window", url).Start()
	return nil
}

func (b *BraveBrowser) OpenInIncognito(url string) error {
	exec.Command("brave-browser", "--incognito", url).Start()
	return nil
}

func NewBrowser(n string) Browser {
	isBrave := strings.Contains("brave-browser", strings.ToLower(n))

	if isBrave {
		return &BraveBrowser{
			binaryPath: n,
		}
	}

	return &BraveBrowser{
		binaryPath: "brave-browser",
	}
}
