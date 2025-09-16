package main

import (
	"fmt"
)

func main() {
	fmt.Printf("111111111\n")
	defer fmt.Printf("defer 1\n")
	defer fmt.Printf("defer 2\n")
	defer fmt.Printf("defer 3\n")
	fmt.Printf("1111111112\n")
	fmt.Printf("1111111113\n")
}
