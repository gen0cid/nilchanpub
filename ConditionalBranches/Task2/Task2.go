package main

import "fmt"

func main() {
	age := 11
	if age >= 18 {
		fmt.Println("Я продам тебе пиво.")
	} else if age < 12 {
		fmt.Println("Ты чего малой совсем берега попутал пиво покупать?!")
	} else {
		fmt.Println("Алкоголь только от 18-ти лет")
	}
}
