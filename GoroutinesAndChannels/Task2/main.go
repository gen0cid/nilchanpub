package main

import (
	"fmt"
	"math/rand"
	"time"
)

func foo(transferPoint chan int, n int) {
	b := rand.Intn(150)
	fmt.Println(b)

	transferPoint <- b
	fmt.Println("Уголь передал рабочий", n)
}

func main() {
	sum := 0

	transferPoint := make(chan int, 4)

	go foo(transferPoint, 1)
	go foo(transferPoint, 2)
	go foo(transferPoint, 3)
	go foo(transferPoint, 4)

	sum += <-transferPoint
	fmt.Println("Главный рабочий забрал уголь ")
	time.Sleep(1 * time.Second)

	sum += <-transferPoint
	fmt.Println("Главный рабочий забрал уголь ")
	time.Sleep(1 * time.Second)

	sum += <-transferPoint
	fmt.Println("Главный рабочий забрал уголь ")
	time.Sleep(1 * time.Second)

	sum += <-transferPoint
	fmt.Println("Главный рабочий забрал уголь ")
	time.Sleep(1 * time.Second)

	fmt.Println(sum)
}
