package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	go printHi()
	fmt.Println("Hello, World!2")
	wg := sync.WaitGroup{}
	for range 10 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			printGoogle()
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(t))
}

func printGoogle() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)

}

func printHi() {
	fmt.Println("Hello, goroutine!")
}
