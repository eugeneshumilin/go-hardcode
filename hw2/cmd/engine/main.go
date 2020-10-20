package main

import (
	"bufio"
	"fmt"
	"go-hardcode/hw2/pkg/spider"
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

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Найти: ")
		req, _ := reader.ReadString('\n')
		req = strings.Replace(req, "\n", "", -1)

		res := make(map[string]string)

		for k, v := range links {
			if strings.Contains(k, req) || strings.Contains(v, req) {
				res[k] = v
			}
		}

		for k, v := range res {
			fmt.Printf("Link: %v, Title: %v \n", k, v)
		}
	}

}
