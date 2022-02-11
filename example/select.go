package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	channel := make(chan bool)
	rand.Seed(time.Now().Unix()) // set up seed for random

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second) // simulate working
		channel <- true
	}()

	timeout := time.Tick(3 * time.Second) // timeout = 3 seconds

	select {
	case <-timeout:
		fmt.Println("Timeout!")
	case <-channel:
		fmt.Println("Work done!")
	}
}