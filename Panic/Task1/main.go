package main

import "fmt"

func testDivideByZero() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была паника:", p)
		}
	}()

	a := 4
	b := 0
	c := a / b
	fmt.Println(c)

}

func testNilMap() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была паника:", p)
		}
	}()
	var nilmap map[int]string
	nilmap[21] = "облаачно"

}

func testOutOfBounds() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Была паника:", p)
		}
	}()

	slice := []int{2, 5, 8}
	slice[3] = 34
}

func main() {
	testDivideByZero()
	testNilMap()
	testOutOfBounds()
}
