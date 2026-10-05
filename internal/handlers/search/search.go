package search

import "github.com/WardenDigital/g-rofi/internal/rofi"

func Search(b string, e string) error {
	browser := NewBrowser(b)
	engine := NewSearchEngine(e)

	query, err := rofi.Prompt("Enter search query:")

	if err != nil {
		return err
	}

	return performSearch(browser, engine, string(query))
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
