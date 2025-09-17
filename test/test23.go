package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("i=%d\n", i)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello1")
	say("hello2")
	say("hello3")
}
