package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	code := make(chan int)
	wg := sync.WaitGroup{}
	for range 10 {
		wg.Add(1)
		go func() {
			printGoogle(code)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(code)
	}()
	for res := range code {
		fmt.Println(res)
	}

}

func printGoogle(code chan int) {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	code <- resp.StatusCode
}

func printHi() {
	fmt.Println("Hello, goroutine!")
}
