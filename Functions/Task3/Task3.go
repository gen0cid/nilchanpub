package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	a = 2
	b = -4
	c = -6
	square(a, b, c)
}
func square(a, b, c float64) {
	d := b*b - 4*a*c
	if d < 0 {
		fmt.Println("Корней нет, тк D<0")
		return
	}
	if d == 0 {
		x := -b / (2 * a)
		fmt.Println("Один корень, тк D=0")
		fmt.Println("x:", x)
		return
	}
	x1 := (-b - math.Sqrt(d)) / (2 * a)
	x2 := (-b + math.Sqrt(d)) / (2 * a)
	fmt.Println("x1:", x1)
	fmt.Println("x2:", x2)
	return
}
