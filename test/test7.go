package main

import "fmt"

type A struct {
	a int8
	b int
}

func main() {
	v := A{}
	v.a = 1
	v.b = 2

	fmt.Printf("fkjsdfkjsakf v=[%v][%v][%v][%v]\n", v, &v, &v.a, &v.b)
}
