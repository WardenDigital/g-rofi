package search

import (
	"strings"
	"testing"
)

func TestNewBrowser(t *testing.T) {
	if _, ok := NewBrowser("librewolf").(*LibrewolfBrowser); !ok {
		t.Error("librewolf should match LibrewolfBrowser")
	}
	if _, ok := NewBrowser("brave-browser").(*BraveBrowser); !ok {
		t.Error("brave-browser should match BraveBrowser")
	}
	if _, ok := NewBrowser("librewolf-nightly").(*LibrewolfBrowser); !ok {
		t.Error("librewolf-nightly should match LibrewolfBrowser")
	}
	if _, ok := NewBrowser("x-unknown-y").(*BraveBrowser); !ok {
		t.Error("unknown should fall back to BraveBrowser")
	}
	if _, ok := NewBrowser("Brave-Browser").(*BraveBrowser); !ok {
		t.Error("case-insensitive brave")
	}
}

func TestSearchURLEncoding(t *testing.T) {
	u, err := (&GoogleSearch{}).Search("two words & more")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(u, "two+words+%26+more") {
		t.Errorf("google query not encoded: %s", u)
	}
	u2, err := (&StartPageSearch{}).Search("a b")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(u2, "q=a+b") {
		t.Errorf("startpage query not encoded: %s", u2)
	}
}
