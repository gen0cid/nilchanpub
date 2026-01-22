package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func(ch chan int) {
		res := 100 * (1 + rand.IntN(10))
		time.Sleep(time.Duration(res) * time.Millisecond)
		fmt.Printf("Сгенерированное количество времени 1 горутины: %d\n", res)
		ch <- res
	}(chan1)

	go func(ch chan int) {
		res := 100 * rand.IntN(11)
		time.Sleep(time.Duration(res) * time.Millisecond)
		fmt.Printf("Сгенерированное количество времени 2 горутины: %d\n", res)
		ch <- res
	}(chan2)

	time.Sleep(500 * time.Millisecond)

	select {
	case I1 := <-chan1:
		fmt.Println("Отработала секция select, чтение канала chan1:", I1)

	case I2 := <-chan2:
		fmt.Println("Отработала секция select, чтение канала chan2:", I2)

	default:
		fmt.Println("Отработала секция select default")
	}
}
