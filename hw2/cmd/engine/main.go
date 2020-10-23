package main

import (
	"bufio"
	"fmt"
	"go-hardcode/hw2/pkg/spider"
	"log"
	"os"
	"strings"
)

func main() {
	urls := []string{"https://go.dev/", "https://dave.cheney.net/"}

	links := make(map[string]string)

	for _, v := range urls {
		l, _ := spider.Scan(v, 2)
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
		req = strings.Replace(req, "\n", "", -1)
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
