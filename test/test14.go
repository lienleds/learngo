package main

import "fmt"

type A struct {
	a int
	b int
}

func ff(a A) {
	a.a = 1
	a.b = 2
	fmt.Printf("ff=[%v][%v]\n", a, &a.a)
}

func fff(a *A) {
	a.a = 3
	a.b = 4
	fmt.Printf("fff=[%v][%v]\n", a, &a.a)
}

func main() {
	a := A{1, 2}
	fmt.Printf("a1=[%v][%v]\n", a, &a.a)
	ff(a)
	fmt.Printf("a2=[%v][%v]\n", a, &a.a)
	fff(&a)
	fmt.Printf("a3=[%v][%v]\n", a, &a.a)
}
