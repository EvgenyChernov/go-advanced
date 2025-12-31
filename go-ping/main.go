package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func pingUrl(url string, respCh chan int, errCh chan error) {

	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()
	respCh <- resp.StatusCode

}

func main() {
	path := flag.String("file", "url.txt", "path to url file")
	flag.Parse()

	file, err := os.ReadFile(*path)

	if err != nil {
		panic(err.Error())
	}
	urlSlice := strings.Split(string(file), "\n")
	respCh := make(chan int)
	errCh := make(chan error)
	for _, url := range urlSlice {
		go pingUrl(url, respCh, errCh)
	}
	for range urlSlice {
		select {
		case res := <-respCh:
			fmt.Println(res)
		case err := <-errCh:
			fmt.Println(err)
		}
	}
}
