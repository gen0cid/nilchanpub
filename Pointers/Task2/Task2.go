package main

import "fmt"

func main() {
	number := 45
	text := "Привет"
	boolean := false
	drob := 34.546

	fmt.Println("Переменненые до вызова функций ")
	fmt.Println()

	fmt.Println("numder:", number)
	fmt.Println("text:", text)
	fmt.Println("boolean:", boolean)
	fmt.Println("drob:", drob)

	fmt.Println()
	fmt.Println("Переменненые после вызова функций ")
	fmt.Println()

	numberPtr := &number
	textPtr := &text
	booleanPtr := &boolean
	drobPtr := &drob

	foo(numberPtr)
	boo(textPtr)
	loo(booleanPtr)
	moo(drobPtr)

	fmt.Println("numder:", number)
	fmt.Println("text:", text)
	fmt.Println("boolean:", boolean)
	fmt.Println("drob:", drob)

}

func foo(n *int) {
	*n = 23
}

func boo(n *string) {
	*n = "Как дела?"
}

func loo(n *bool) {
	*n = true
}

func moo(n *float64) {
	*n = 56.87
}
