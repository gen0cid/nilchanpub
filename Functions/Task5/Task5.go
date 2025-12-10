package main

import "fmt"

var number float64 = 54

func main() {
	fmt.Println("number1:")
	fmt.Println(number)
	Go()

	fmt.Println("number2:")
	fmt.Println(number)
	boo()

	fmt.Println("number3:")
	fmt.Println(number)
	sii()

	fmt.Println("number4:")
	fmt.Println(number)
}
func Go() {
	number -= 5
}

func boo() {
	number /= 25
}
func sii() {
	number += 47
}
