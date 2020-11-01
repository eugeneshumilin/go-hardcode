// Package index индексирует документы и хранит в объекте
package index

import "testing"

func TestIndex_FillStorage(t *testing.T) {
	ind := New()
	data := map[string]string{
		"https://go.dev/":          "go dev",
		"https://dave.cheney.net/": "Dave Cheney",
	}
	ind.FillStorage(data)
	got := len(ind.Storage)
	want := 2
	if got != want {
		t.Fatalf("получили %d, ожидалось %d", got, want)
	}
}

func TestIndex_Search(t *testing.T) {
	ind := New()
	data := map[string]string{
		"https://go.dev/":          "go dev",
		"https://dave.cheney.net/": "Dave Cheney",
	}
	ind.FillStorage(data)
	ind.Search("dave")
	got := ind.Search("dave")[0]
	want := "https://dave.cheney.net/ - Dave Cheney \n"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}
