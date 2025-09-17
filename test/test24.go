package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func f1(i int) {
	mu.Lock()
	fmt.Printf("====> f1 [%d]\n", i)
	time.Sleep(1000 * time.Millisecond)
	mu.Unlock()
}

func f2() {
	for {
		fmt.Printf("====> f2\n")
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	fmt.Printf("====> main 1\n")
	go f1(1)
	go f1(2)
	f1(1)
	fmt.Printf("====> main 2\n")
	// go f2()
	// fmt.Printf("====> main 3\n")
	// f2()
	// fmt.Printf("====> main 4\n")
}
