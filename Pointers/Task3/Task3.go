package main

import (
	"fmt"
)

func main() {
	number := 345
	text := "Привет, как дела?"
	boolean := false
	drob := 345.34

	numberPtr1 := &number
	textPtr1 := &text
	booleanPtr1 := &boolean
	drobPtr1 := &drob

	var numberPtr2 *int
	var textPtr2 *string
	var booleanPtr2 *bool
	var drobPtr2 *float64

	proverkaNumberPtr1(numberPtr1)
	proverkatextPtr1(textPtr1)
	proverkabooleanPtr1(booleanPtr1)
	proverkadrobPtr1(drobPtr1)

	proverkanumberPtr2(numberPtr2)
	proverkatextPtr2(textPtr2)
	proverkabooleanPtr2(booleanPtr2)
	proverkadrobPtr2(drobPtr2)
}

func proverkaNumberPtr1(n *int) {
	if n == nil {
		fmt.Println("Указатель nil", n)
		fmt.Println()
		return
	}
	fmt.Println("Указатель не nil,", n, " указывает на переменную со значением:", *n)
	fmt.Println()
}

func proverkatextPtr1(n *string) {
	if n == nil {
		fmt.Println("Указатель nil", n)
		fmt.Println()
		return
	}
	fmt.Println("Указатель не nil,", n, " указывает на переменную со значением:", *n)
	fmt.Println()
}

func proverkabooleanPtr1(n *bool) {
	if n == nil {
		fmt.Println("Указатель nil", n)
		fmt.Println()
		return
	}
	fmt.Println("Указатель не nil,", n, " указывает на переменную со значением:", *n)
	fmt.Println()
}

func proverkadrobPtr1(n *float64) {
	if n == nil {
		fmt.Println("Указатель nil", n)
		fmt.Println()
		return
	}
	fmt.Println("Указатель не nil,", n, " указывает на переменную со значением:", *n)
	fmt.Println()
}

func proverkanumberPtr2(n *int) {
	if n == nil {
		fmt.Println("Указатель nil", n)
		fmt.Println()
		return
	}
	fmt.Println("Указатель не nil,", n, " указывает на переменную со значением:", *n)
	fmt.Println()
}

func proverkatextPtr2(n *string) {
	if n == nil {
		fmt.Println("Указатель nil", n)
		fmt.Println()
		return
	}
	fmt.Println("Указатель не nil,", n, " указывает на переменную со значением:", *n)
	fmt.Println()
}

func proverkabooleanPtr2(n *bool) {
	if n == nil {
		fmt.Println("Указатель nil", n)
		fmt.Println()
		return
	}

	fmt.Println("Указатель не nil,", n, " указывает на переменную со значением:", *n)
	fmt.Println()
}

func proverkadrobPtr2(n *float64) {
	if n == nil {
		fmt.Println("Указатель nil", n)
		fmt.Println()
		return
	}
	fmt.Println("Указатель не nil,", n, " указывает на переменную со значением:", *n)
	fmt.Println()
}
