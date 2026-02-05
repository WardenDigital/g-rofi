package search

type SearchEngine interface {
	Search(query string) (string, error)
}

type GoogleSearch struct{}

func (g *GoogleSearch) Search(query string) (string, error) {
	return "https://www.google.com/search?q=" + query, nil
}
