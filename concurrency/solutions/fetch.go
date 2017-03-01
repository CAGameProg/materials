package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v\n", err)
		return
	}
	resp.Body.Close()
	ch <- fmt.Sprintf("%.2fs %s\n", time.Since(start).Seconds(), url)
}

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Print(<-ch)
	}
}
