package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	score := 0
	fmt.Println("Продвинутая симуляция игры Geometry Dash")
	fmt.Println("--------------------")

	for {

		fmt.Println("Я подползю к шипам!")
		fmt.Println("* /\\")
		fmt.Println()
		fmt.Println("Я делаю прыжок!")
		fmt.Println("*")
		fmt.Println("/\\")
		fmt.Println()

		//0 1 2 3 4 5 6 7
		if rand.Intn(8) == 1 {
			fmt.Println("Я упал в шипы!")
			fmt.Println("/Х\\")
			fmt.Println("=========================")
			break
		} else {
			fmt.Println("Я перепрыгнул через шипы!")
			fmt.Println("/\\ *")
			score += 5
			fmt.Println("--------------------")
			fmt.Println("Ваш счет:", score)
			fmt.Println()
			time.Sleep(500 * time.Millisecond)

		}
	}
	fmt.Println("Game over!")
	fmt.Println("Итоговый счет:", score)
}
