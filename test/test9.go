package main

import "fmt"

func main() {
	var sli [3]int
	sli[0] = 0x1
	sli[1] = 0x2
	sli[2] = 0x3
	fmt.Printf("lkdsjflsa[%v] %v\n", sli, cap(sli))
	a := sli[0:1]
	fmt.Printf("lkdsjflsa[%v]\n", a)
	a[0] = 0x9
	fmt.Printf("lkdsjflsa[%v][%v]\n", sli, a)

	var ss []int
	fmt.Printf("ss=[%v]\n", ss)

	if ss == nil {
		fmt.Printf("Nil Nil Nil Nil Nil \n")
	} else {
		fmt.Printf("Not Nil\n")
	}
}
