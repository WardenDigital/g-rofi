package search

import "os/exec"

type Browser interface {
	Open(url string) error
	OpenInNewWindow(url string) error
	OpenInIncognito(url string) error
}

type BraveBrowser struct{}

func (b *BraveBrowser) Open(url string) error {
	exec.Command("brave-browser", url).Start()

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
