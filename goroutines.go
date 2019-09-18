package main

import (
	"fmt"
	"time"
)

func longTask() {
	fmt.Println("Starting long task")
	time.Sleep(3 * time.Second)
	fmt.Println("Long task finished")
}

func main() {
	go longTask()
	go longTask()
	go longTask()
	time.Sleep(4 * time.Second)
}
