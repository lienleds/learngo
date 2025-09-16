package main

import "fmt"

type A struct {
	a int
	b int
}

func main() {

	m := make(map[string]A)
	m["a"] = A{1, 2}
	m["b"] = A{3, 4}

	fmt.Printf("m=[%v]\n", m)

	n, ok := m["c"]
	fmt.Printf("n=[%v]\n", n)
	fmt.Printf("ok=[%v]\n", ok)

	n1, ok1 := m["a"]
	fmt.Printf("n=[%v]\n", n1)
	fmt.Printf("ok=[%v]\n", ok1)

	delete(m, "b")
	fmt.Printf("m=[%v]\n", m)

}
