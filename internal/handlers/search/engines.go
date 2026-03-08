package search

type SearchEngine interface {
	Search(query string) (string, error)
}

type GoogleSearch struct{}

func (g *GoogleSearch) Search(query string) (string, error) {
	return "https://www.google.com/search?q=" + query, nil
}

type StartPageSearch struct{}

func (s *StartPageSearch) Search(query string) (string, error) {
	return "https://www.startpage.com/sp/search?q=" + query, nil
}

func NewSearchEngine(name string) SearchEngine {
	switch name {
	case "google":
		return &GoogleSearch{}
	case "startpage":
		return &StartPageSearch{}
	}

	return &GoogleSearch{}
}
