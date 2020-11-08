// Package btree для построения бинарного дерева для сущности model.Document
package btree

import (
	"fmt"
	"go-hardcode/hw5/pkg/model"
)

// Tree - Двоичное дерево поиска
type Tree struct {
	root *Element
}

// Element - элемент дерева
type Element struct {
	left, right *Element
	Value       *model.Document
}

// Insert - вставка элемента в дерево
func (t *Tree) Insert(d *model.Document) {
	e := &Element{Value: d}
	if t.root == nil {
		t.root = e
		return
	}
	insert(t.root, e)
}

// insert рекурсивно вставляет элемент в нужный уровень дерева.
func insert(node, new *Element) {
	if new.Value.Id < node.Value.Id {
		if node.left == nil {
			node.left = new
			return
		}
		insert(node.left, new)
	}
	if new.Value.Id >= node.Value.Id {
		if node.right == nil {
			node.right = new
			return
		}
		insert(node.right, new)
	}
}

// Search - поиск значения в дереве, возвращает документ и true если найдено
func (t *Tree) Search(id int) (*model.Document, bool) {
	node := t.root

	for {
		if node == nil {
			return &model.Document{}, false
		}

		if node.Value.Id == id {
			return node.Value, true
		}

		if node.Value.Id > id {
			node = node.left
			continue
		}

		node = node.right
	}
}

// строковое представление бинарного дерева
func (t *Tree) String() string {
	ids := []int{}

	t.root.collectIds(&ids)
	return fmt.Sprint(ids)

}

// заполняем слайс айдишниками документов
func (e *Element) collectIds(ids *[]int) {
	if e == nil {
		return
	}

	*ids = append(*ids, e.Value.Id)

	e.left.collectIds(ids)
	e.right.collectIds(ids)
}
