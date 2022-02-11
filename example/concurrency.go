package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Task 2")
		}
	}()

	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Task 1")
	}
}