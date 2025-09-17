package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot sqrt negative number")
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
