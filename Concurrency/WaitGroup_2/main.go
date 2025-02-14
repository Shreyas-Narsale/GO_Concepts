package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var urls = []string{"https://go.dev", "https://google.com", "https://github.com"}

	for _, url := range urls {
		wg.Add(1)
		go GetStatusCode(url)
	}
	wg.Wait()
}

func GetStatusCode(url string) {
	defer wg.Done()
	fmt.Println("Endpoint", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("StatusCode", resp.StatusCode)
	}
}
