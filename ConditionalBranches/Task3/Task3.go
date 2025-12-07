package main

import "fmt"

func main() {
	liters := 2.5

	if liters < 1 || liters > 3 {
		fmt.Println("Ты какой-то странный...")
	} else {
		fmt.Println("А ты знаешь золотую середину!")
	}
}
