package main

import "fmt"

func main() {
	number := 13

	fmt.Println("-------------------------------")

	fmt.Println("number до +10:", number)
	number += 10
	fmt.Println("number после +10:", number)

	fmt.Println("-------------------------------")

	fmt.Println("number до -10:", number)
	number -= 10
	fmt.Println("number после -10:", number)

	fmt.Println("-------------------------------")

	fmt.Println("number до *5:", number)
	number *= 5
	fmt.Println("number после *5:", number)

	fmt.Println("-------------------------------")

	fmt.Println("number до /5:", number)
	number /= 5
	fmt.Println("number после /5:", number)

	fmt.Println("-------------------------------")

	fmt.Println("number до %3:", number)
	number %= 3
	fmt.Println("number после %3:", number)

	fmt.Println("-------------------------------")

	fmt.Println("number до ++ :", number)
	number++
	fmt.Println("number после ++ :", number)

	fmt.Println("-------------------------------")

	fmt.Println("number до -- :", number)
	number--
	fmt.Println("number после -- :", number)

	fmt.Println("-------------------------------")
}
