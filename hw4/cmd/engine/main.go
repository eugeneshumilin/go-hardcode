package main

import (
	"bufio"
	"fmt"
	"go-hardcode/hw4/pkg/index"
	"go-hardcode/hw4/pkg/spider"
	"log"
	"os"
	"strings"
)

func main() {
	urls := []string{"https://go.dev/", "https://dave.cheney.net/"}

	links := make(map[string]string)

	log.Println("Начинается сканирование документов...")
	for _, v := range urls {
		s := spider.ScanBot{}
		l, err := scanLinks(&s, v, 2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for u, t := range l {
			links[u] = t
		}
	}
	log.Println("Сканирование документов закончено")

	// индексируем документы, заполняем ими наш storage и создаем инвертированный индекс
	log.Println("Начинается индексирование документов...")
	ind := index.New()
	ind.FillStorage(links)
	log.Println("Индексирование закончено")

	// работаем с пользовательским вводом
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Найти: ")
		req, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		req = strings.TrimSuffix(req, "\r\n")
		req = strings.TrimSuffix(req, "\n")

		targetWords := strings.Split(req, " ")

		// получаем слайс строк вида "url - title"
		output := []string{}
		for _, word := range targetWords {
			output = append(output, ind.Search(word)...)
		}

		// вывод результата
		for _, doc := range output {
			fmt.Println(doc)
		}
	}
}

type Scanner interface {
	Scan(url string, depth int) (map[string]string, error)
}

func scanLinks(s Scanner, url string, depth int) (map[string]string, error) {
	data, err := s.Scan(url, depth)
	return data, err
}
