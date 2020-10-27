package main

import (
	"bufio"
	"fmt"
	"go-hardcode/hw3/pkg/spider"
	"log"
	"os"
	"strings"
)

func main() {
	urls := []string{"https://go.dev/", "https://dave.cheney.net/"}

	links := make(map[string]string)

	for _, v := range urls {
		s := spider.ScanBot{}
		l, err := scanLinks(&s, v, 2)
		if err != nil {
			fmt.Println("Ошибка...")
			return
		}

		for u, t := range l {
			links[u] = t
		}
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Найти: ")
		req, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		req = strings.TrimSuffix(req, "\r\n")
		req = strings.TrimSuffix(req, "\n")
		res := make(map[string]string)

		for k, v := range links {
			if strings.Contains(strings.ToLower(k), strings.ToLower(req)) || strings.Contains(strings.ToLower(v), strings.ToLower(req)) {
				res[k] = v
			}
		}

		for k, v := range res {
			fmt.Printf("Link: %v, Title: %v \n", k, v)
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
