package main

import "fmt"

func main() {
	greeting("Сергей", 127)
	greeting("Даниил", 23)
	greeting("Иван", 13)

}
func greeting(name string, nomer int) {
	fmt.Println("Здравствуйте,", name)
	fmt.Println(name, ",ведем вас в номер", nomer)
	fmt.Println()
}
