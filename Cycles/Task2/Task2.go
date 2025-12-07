package main

import (
	"fmt"
	"time"
)

func main() {
	score := 0
	fmt.Println("простой симулятор игры Geometry Dash")
	fmt.Println("--------------------------------")
	fmt.Println("ваш счет:", score)

	for i := 1; i <= 5; i++ {
		fmt.Println("*/\\")
		fmt.Println("*")
		fmt.Println("/\\")
		fmt.Println("*/\\*")
		fmt.Println("--------------------------------")
		score += 5
		fmt.Println("ваш счет:", score)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("--------------------------------")
	fmt.Println("Уровень пройден!")
	fmt.Println("Итоговый счет: ", score)
}
