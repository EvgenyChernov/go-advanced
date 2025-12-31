package main

import (
	"fmt"
	"net/http"
)

func main() {
	// code := make(chan int)
	// wg := sync.WaitGroup{}
	// for range 10 {
	// 	wg.Add(1)
	// 	go func() {
	// 		printGoogle(code)
	// 		wg.Done()
	// 	}()
	// }
	// go func() {
	// 	wg.Wait()
	// 	close(code)
	// }()
	// for res := range code {
	// 	fmt.Println(res)
	// }

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 3
	ch := make(chan int, numGoroutines)

	partSize := len(arr) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		go sumPart(arr[start:end], ch)
	}

	totalSum := 0
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-ch
	}

	fmt.Println("Total sum:", totalSum)
}

func sumPart(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
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
