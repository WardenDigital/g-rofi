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
	exec.Command(b.binaryPath, "--new-window", url).Start()
	return nil
}

func (b *BraveBrowser) OpenInIncognito(url string) error {
	exec.Command(b.binaryPath, "--incognito", url).Start()
	return nil
}

type LibrewolfBrowser struct {
	binaryPath string
}

func (b *LibrewolfBrowser) Open(url string) error {
	exec.Command(b.binaryPath, url).Start()

	return nil
}

func (b *LibrewolfBrowser) OpenInNewWindow(url string) error {
	exec.Command(b.binaryPath, "--new-window", url).Start()
	return nil
}

func (b *LibrewolfBrowser) OpenInIncognito(url string) error {
	exec.Command(b.binaryPath, "--private-window", url).Start()
	return nil
}

func NewBrowser(n string) Browser {
	isBrave := strings.Contains("brave-browser", strings.ToLower(n))
	isLibrewolf := strings.Contains("librewolf", strings.ToLower(n))

	switch true {
	case isBrave:
		return &BraveBrowser{
			binaryPath: n,
		}
	case isLibrewolf:
		return &LibrewolfBrowser{
			binaryPath: n,
		}
	}

	return &BraveBrowser{
		binaryPath: n,
	}
}
