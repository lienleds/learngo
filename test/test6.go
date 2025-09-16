package main

import (
	"fmt"
)

// Renamed to main2 to avoid redeclaration error
func main() {
	var i *int
	j := 123

	i = &j
	if i != nil {
		fmt.Printf("aaa :[%v][%v][%v]\n", i, *i, &j)
		j = 456
		fmt.Printf("bbb :[%v][%v][%v]\n", i, *i, &j)
		*i = 789
		fmt.Printf("ccc :[%v][%v][%v]\n", i, *i, &j)
	} else {
		fmt.Printf("dklfjadklsfjkdjs\n")
	}
}
