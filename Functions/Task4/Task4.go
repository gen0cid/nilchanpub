package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64 = 12, 32
	F := linear(a, b)
	fmt.Println("Значение функции:", F)
}
func linear(x, y float64) float64 {
	return 2*math.Pow(x, 3) + 3*y - 5
}
