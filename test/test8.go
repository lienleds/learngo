package main

import "fmt"

func main() {
	var arr [5]int
	idx := 5

	arr[0] = 99
	arr[1] = 98

	fmt.Printf("Array [%v]\n", arr)
	arr[idx] = 0xFA
}
