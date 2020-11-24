// Package btree для построения бинарного дерева для сущности model.Document
package btree

import (
	"go-hardcode/hw5/pkg/model"
	"testing"
)

func seed(t *Tree) {
	t.Insert(&model.Document{Id: 5})
	t.Insert(&model.Document{Id: 1})
	t.Insert(&model.Document{Id: 2})
	t.Insert(&model.Document{Id: 3})
	t.Insert(&model.Document{Id: 80})
	t.Insert(&model.Document{Id: 17})
}

func TestTree_Insert(t *testing.T) {
	tree := Tree{}

	got := tree.String()
	want := "[]"
	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}

	seed(&tree)

	got = tree.String()
	want = "[5 1 2 3 80 17]"
	if got != want {
		t.Fatalf("got %s; want %s", got, want)
	}
}

func TestTree_Search(t *testing.T) {
	tree := Tree{}
	doc := &model.Document{Id: 33}

	tree.Insert(doc)

	seed(&tree)

	got, ok := tree.Search(33)
	if !ok {
		t.Fatalf("got %v; want true", ok)
	}
	want := doc
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}

	_, ok = tree.Search(10001)
	if ok {
		t.Fatalf("got %v; want false", ok)
	}
}
