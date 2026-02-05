package search

import "github.com/WardenDigital/rofi-wrapper/internal/rofi"

func Search() {
	browser := &BraveBrowser{}
	engine := &GoogleSearch{}

	query, err := rofi.Prompt("Enter search query:")

	if err != nil {
		panic(err)
	}

	err = performSearch(browser, engine, string(query))

	if err != nil {
		panic(err)
	}
}

func performSearch(b Browser, e SearchEngine, q string) error {
	url, err := e.Search(q)
	if err != nil {
		return err
	}

	err = b.OpenInNewWindow(url)

	if err != nil {
		return err
	}

	return nil
}
