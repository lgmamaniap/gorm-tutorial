package main

import (
	"fmt"
	"time"
)

// Equivalent to async await in JS
// Channels are a powerful feature in Go that allow goroutines to communicate with each other.

func longTask(result chan string) {
	time.Sleep(2 * time.Second)
	result <- "Task completed!"
}

func main() {
	channel := make(chan string)
	go longTask(channel)
	fmt.Println("Waiting for task to complete...")
	result := <-channel
	fmt.Println(result)
}
