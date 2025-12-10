package main

import "fmt"

func main() {
	fmt.Println("Я main и я начался")
	defer func() {
		fmt.Println("привет я defer1 и я закончился")
	}()
	defer func() {
		fmt.Println("привет я defer2 и я закончился")
	}()
	defer func() {
		fmt.Println("привет я defer3 и я закончился")
	}()
	fmt.Println("Я main и я закончился")

}
