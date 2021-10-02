package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
type Page struct {
	URL string
	Size int
}
func main() {
	pages := make(chan Page)
	var urls [3]string
	urls = [3]string{
		"https://golang.org/src",
		"https://golang.org/doc",
		"https://golang.org/pkg",
	}
	for _, url := range urls {
		go responseSize(pages, url)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Println("Getting", page.URL, "Size: ",  page.Size, "bytes.")
	}
}

func responseSize(channel chan Page, url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}
