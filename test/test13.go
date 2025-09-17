package main

import "fmt"

type A struct {
	a int
	b int
}

func (a A) sum() int {
	return a.a + a.b
}

func main() {
	a := A{1, 2}
	b := a.sum()
	fmt.Printf("main = [%v]\n", b)
}
