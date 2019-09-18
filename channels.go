package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTask(channel chan int) {
	delay := rand.Intn(5)
	fmt.Println("Starting long task")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("Long task finished")
	channel <- delay
}

func main() {
	rand.Seed(time.Now().Unix())

	channel := make(chan int)
	go longTask(channel) // call longTask as a goroutine
	fmt.Println("Took", <-channel, "seconds")
}
