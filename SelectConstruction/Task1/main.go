package main

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)
	flCh := make(chan float64)

	go func(intCh chan int) {
		i := 0
		for {
			intCh <- i
			i++
			time.Sleep(300 * time.Millisecond)
		}
	}(intCh)

	go func(strCh chan string) {
		for {
			strCh <- "привет"
			time.Sleep(500 * time.Millisecond)
		}
	}(strCh)

	go func(flCh chan float64) {
		i := 1.0
		for {
			flCh <- i
			i /= 2
			time.Sleep(1 * time.Second)
		}
	}(flCh)
	for {
		select {
		case I := <-intCh:
			fmt.Printf("Читаю сообщение из канала intCh: %d\n", I)
		case S := <-strCh:
			fmt.Printf("Читаю сообщение из канала strCh: %s\n", S)
		case F := <-flCh:
			fmt.Printf("Читаю сообщение из канала flCh: %f\n", F)
		}
	}

}
