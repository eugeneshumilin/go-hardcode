// Package index индексирует документы и хранит в объекте
package index

import "testing"

func TestIndex_Search(t *testing.T) {
	ind := New()
	data := map[string]string{
		"https://go.dev/":          "go dev",
		"https://dave.cheney.net/": "Dave Cheney",
		"https://go.com/":          "go go go",
	}
	ind.FillStorage(data)

	found := ind.Search("go")
	got := len(found)
	want := 2
	if got != want {
		t.Fatalf("len(found) = %d; want %d", got, want)
	}
}
