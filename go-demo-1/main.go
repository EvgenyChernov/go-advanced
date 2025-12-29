package main

import (
	"fmt"
	"net/http"
)

func main() {
	code := make(chan int)
	go printGoogle(code)
	fmt.Println(<-code)
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
