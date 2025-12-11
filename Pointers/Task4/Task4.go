package main

import (
	"fmt"
)

func main() {
	var size *float64
	fmt.Println("До операции:", size)
	changeSize(size)
	fmt.Println("После операции:", size)

}
func changeSize(n *float64) {
	if n == nil {
		fmt.Println("Операция отменена, тк указатель указывает на nil")
		fmt.Println()
		return
	}
	*n += 1

}
