package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, curr := 0, 1
	fmt.Printf("fibonacci [%v][%v]\n", prev, curr)

	return func() int {
		a := 1
		a++
		temp := prev
		prev = curr
		curr = temp + curr
		fmt.Printf("fibonacci1 [%v][%v][%v][%v]\n", a, temp, prev, curr)
		return prev
	}
}

func main() {
	fmt.Printf("main 1\n")

	f := fibonacci()
	fmt.Printf("main 1\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("main 2\n")
		fmt.Println(f())
		fmt.Printf("main 3\n")
	}
	fmt.Printf("main 4\n")
}
