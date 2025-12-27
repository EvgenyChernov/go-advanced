package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	go printHi()
	fmt.Println("Hello, World!2")

	for i := 0; i < 10; i++ {
		go printGoogle()
	}

	time.Sleep(2 * time.Second)
}

func printGoogle() {
	t := time.Now()
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(time.Since(t))
}

func printHi() {
	fmt.Println("Hello, goroutine!")
}
