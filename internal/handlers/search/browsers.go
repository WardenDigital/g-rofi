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
	name := strings.ToLower(n)
	isBrave := strings.Contains(name, "brave-browser")
	isLibrewolf := strings.Contains(name, "librewolf")

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
