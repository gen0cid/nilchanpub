package Y

import "fmt"

var PublicYGlobal int = 100

func PublicFunc() {
	fmt.Println("я публичная функция")

}

type PublicYStruct struct{}

func (n PublicYStruct) PublicMethod() {
	fmt.Println("Я публичный метод")

}
func VivodYprivat() {
	z := privateYGlobal
	fmt.Println("privateYGlobal", z)
}
