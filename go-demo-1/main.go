package main

import (
	"fmt"
	"time"
)

func main() {

	go printHi()
	fmt.Println("Hello, World!2")
	time.Sleep(1 * time.Second)
}

func printHi() {
	fmt.Println("Hello, goroutine!")
}
