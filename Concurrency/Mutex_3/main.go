package main

import (
	"fmt"
	"net/http"
	"sync"
)

var RechableUrls []string

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	var urls = []string{"https://go.dev", "https://google.com", "https://github.com", "https://shree.com"}

	for _, url := range urls {
		wg.Add(1)
		go GetStatusCode(url)
	}
	wg.Wait()
	fmt.Println("Rechable Urls", RechableUrls)
}

func GetStatusCode(url string) {
	defer wg.Done()
	fmt.Println("Endpoint", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		mu.Lock()
		RechableUrls = append(RechableUrls, url)
		mu.Unlock()
		fmt.Println("StatusCode", resp.StatusCode)
	}
}
