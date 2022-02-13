package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	i := &SafeNumber{}
	go func() {
		i.Set(5)
	}()

	fmt.Println(i.Get())

	time.Sleep(1 * time.Second)
	fmt.Println(i.Get())
}

type SafeNumber struct {
	val int
	m   sync.Mutex
}

func (i *SafeNumber) Get() int {
	i.m.Lock()
	defer i.m.Unlock()

	return i.val
}

func (i *SafeNumber) Set(val int) {
	i.m.Lock()
	defer i.m.Unlock()
	
	i.val = val
}