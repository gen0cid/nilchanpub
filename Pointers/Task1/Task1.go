package main

import "fmt"

func main() {
	var integer int = 15
	var str = "Привет, как дела?"
	var drob = 2345.234
	var boolean bool = true

	integerPtr := &integer
	strPtr := &str
	drobPtr := &drob
	booleanPtr := &boolean

	fmt.Println("integerPtr:", integerPtr)
	fmt.Println("strPtr:", strPtr)
	fmt.Println("drobPtr:", drobPtr)
	fmt.Println("booleanPtr:", booleanPtr)

	fmt.Println("Разыменованный указатель integerPtr:", *integerPtr)
	fmt.Println("Разыменованный указатель strPtr:", *strPtr)
	fmt.Println("Разыменованный указатель drobPtr:", *drobPtr)
	fmt.Println("Разыменованный указатель booleanPtr:", *booleanPtr)
}
