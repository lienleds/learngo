package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello 1111\n")
	c := make(chan int, 3)
	go func(c chan int) {
		fmt.Printf("Hello 22222\n")
		time.Sleep(time.Second * 2)
		fmt.Printf("Hello 3333\n")
		//c <- 42
		//c <- 43
		///c <- 44
		fmt.Printf("Hello 4444\n")
		//close(c)
	}(c)

	fmt.Printf("Hello 5555\n")
	ret := <-c
	fmt.Printf("Hello 666 %d %d\n", ret, len(c))
	ret1 := <-c
	fmt.Printf("Hello 666 %d\n", ret1)

	ret2 := <-c
	fmt.Printf("Hello 666 %d\n", ret2)

	ret3 := <-c
	fmt.Printf("Hello 666 %d\n", ret3)
	go func(c chan int) {
		fmt.Printf("Hello 777\n")
		time.Sleep(time.Second * 2)
		fmt.Printf("Hello 888\n")
		c <- 42
		c <- 43
		c <- 44
		fmt.Printf("Hello 999\n")
		close(c)
	}(c)

	time.Sleep(time.Second * 3)
	ret4 := <-c
	fmt.Printf("Hello 666 %d\n", ret4)
}
