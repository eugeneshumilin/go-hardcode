// Package index индексирует документы и хранит в объекте
package index

import (
	"fmt"
	"sort"
	"strings"
)

type Index struct {
	Storage       []Document
	InvertedIndex map[string][]int
}

type Document struct {
	Id    int
	Url   string
	Title string
}

// функция-конструктор нового объекта Индекса
func New() *Index {
	ind := Index{
		Storage:       []Document{},
		InvertedIndex: map[string][]int{},
	}
	return &ind
}

// Заполняем Index.Storage документами типа Document, сортируем их по id, строим инвертированный индекс
func (i *Index) FillStorage(data map[string]string) {
	id := 1
	for u, t := range data {
		item := Document{
			Id:    id,
			Url:   u,
			Title: t,
		}
		i.Storage = append(i.Storage, item)
		id++
	}

	sort.Slice(i.Storage, func(a, b int) bool { return i.Storage[a].Id < i.Storage[b].Id })

	i.createInvertedIndex()
}

// Создаем инвертированный индекс
func (i *Index) createInvertedIndex() {
	for _, doc := range i.Storage {
		for _, token := range tokens(doc.Title) {
			if !exists(i.InvertedIndex[token], doc.Id) {
				i.InvertedIndex[token] = append(i.InvertedIndex[token], doc.Id)
			}
		}
	}
}

// Поиск возвращает слайс строк вида "url - title"
func (i *Index) Search(request string) []string {
	ids := i.InvertedIndex[strings.ToLower(request)]
	result := []string{}

	for _, id := range ids {
		index := sort.Search(len(i.Storage), func(index int) bool { return i.Storage[index].Id >= id })
		if index < len(i.Storage) && i.Storage[index].Id == id {
			doc := i.Storage[index]
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
