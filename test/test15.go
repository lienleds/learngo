package main

import "fmt"

type IA interface {
	sum() int
}

type A struct {
	a int
	b int
}

func (a A) sum() int {
	return a.a + a.b
}

func main() {
	var ia IA
	ia = A{1, 2}

	b := ia.sum()
	fmt.Printf("main = [%v]\n", b)
}
