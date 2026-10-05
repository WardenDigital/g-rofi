package search

import (
	"net/url"
	"strings"
)

type SearchEngine interface {
	Search(query string) (string, error)
}

type GoogleSearch struct{}

func (g *GoogleSearch) Search(query string) (string, error) {
	return "https://www.google.com/search?q=" + url.QueryEscape(query), nil
}

type StartPageSearch struct{}

func (s *StartPageSearch) Search(query string) (string, error) {
	return "https://www.startpage.com/sp/search?q=" + url.QueryEscape(query), nil
}

func NewSearchEngine(name string) SearchEngine {
	switch strings.ToLower(name) {
	case "google":
		return &GoogleSearch{}
	case "startpage":
		return &StartPageSearch{}
	}

	return &GoogleSearch{}
}
