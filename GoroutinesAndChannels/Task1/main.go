package main

import (
	"fmt"
	"time"
)

func gorut(n int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Я горутина %d, делаю вывод на экран в %d раз\n", n, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	for i := 1; i <= 3; i++ {
		go gorut(i)
	}
}
