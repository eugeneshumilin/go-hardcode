// Package index индексирует документы и хранит в объекте
package index

import (
	"fmt"
	"go-hardcode/hw5/pkg/btree"
	"go-hardcode/hw5/pkg/model"
	"strings"
)

type Index struct {
	Storage       btree.Tree
	InvertedIndex map[string][]int
}

// функция-конструктор нового объекта Индекса
func New() *Index {
	ind := Index{
		Storage:       btree.Tree{},
		InvertedIndex: map[string][]int{},
	}
	return &ind
}

// Заполняем Index.Storage документами типа Document, сортируем их по id, строим инвертированный индекс
func (i *Index) FillStorage(data map[string]string) {
	id := 1
	for u, t := range data {
		item := model.Document{
			Id:    id,
			Url:   u,
			Title: t,
		}
		i.Storage.Insert(&item)

		// Создаем инвертированный индекс
		for _, token := range tokens(item.Title) {
			if !exists(i.InvertedIndex[token], item.Id) {
				i.InvertedIndex[token] = append(i.InvertedIndex[token], item.Id)
			}
		}

		id++
	}
}

// Поиск возвращает слайс строк вида "url - title"
func (i *Index) Search(request string) []string {
	ids := i.InvertedIndex[strings.ToLower(request)]
	result := []string{}

	for _, id := range ids {
		if doc, ok := i.Storage.Search(id); ok {
			result = append(result, fmt.Sprintf("%s - %s \n", doc.Url, doc.Title))
		}
	}

	return result
}

// Парсим строку на токены и возвращаем слайс токенов
func tokens(s string) []string {
	words := strings.Split(s, " ")
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

// Проверка наличия элемента в массиве
func exists(ids []int, item int) bool {
	for _, id := range ids {
		if id == item {
			return true
		}
	}
	return false
}
